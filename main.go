package main

import (
	"fly/connection"
	"fly/handler"
	"fly/note"
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

	noteRepo := note.NewRepository(db)
	noteService := note.NewService(noteRepo)
	noteHandler := handler.NewHandlerNote(noteService)

	router := gin.Default()
	router.Use(CorsMiddleware())
	router.POST("/note/create", noteHandler.Create)
	router.GET("/notes", noteHandler.GetAll)
	router.DELETE("/note/reset", noteHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	domain := fmt.Sprintf(":%v", port)
	if err := router.Run(domain); err != nil {
		log.Fatal(err)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
