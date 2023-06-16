package main

import (
	"fly/connection"
	"fly/handler"
	"fly/user"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := connection.NewConnection(".env")
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := handler.NewHandlerUser(userService)

	router := gin.Default()
	router.POST("/user/create", userHandler.Create)
	router.GET("/user/:email", userHandler.GetByEmail)
	router.GET("/user/all", userHandler.GetAll)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	domain := fmt.Sprintf(":%v", port)
	if err := router.Run(domain); err != nil {
		log.Fatal(err)
	}
}
