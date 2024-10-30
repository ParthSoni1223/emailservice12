package main

import (
	"email-service/config"
	"email-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	router := gin.Default()
	router.POST("/send-email", handlers.SendEmail)
	router.Run(":8080")
}
