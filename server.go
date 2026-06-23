package main

import (
	"log"

	"example.com/ticket/controller"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func init() {
		err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
func main() {
	server := gin.Default()
	server.POST("/sms/send", controller.SendSmsHandler)
	server.Run()
}