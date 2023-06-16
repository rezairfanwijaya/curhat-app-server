package main

import (
	"fly/connection"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := connection.NewConnection(".env")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/", home)
	router.GET("/users", users)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	domain := fmt.Sprintf(":%v", port)
	if err := router.Run(domain); err != nil {
		log.Fatal(err)
	}
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"code":    http.StatusOK,
		"message": "hello world",
	})
}

func users(c *gin.Context) {
	user := []struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Role  string `json:"role"`
	}{
		{
			ID:    1,
			Email: "user1@gmail.com",
			Role:  "user",
		}, {
			ID:    2,
			Email: "user2@gmail.com",
			Role:  "user",
		}, {
			ID:    3,
			Email: "admin@admin.com",
			Role:  "admin",
		},
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"code":   http.StatusOK,
		"data":   user,
	})

}
