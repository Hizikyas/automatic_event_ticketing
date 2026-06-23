package controller

import (
	"net/http"

	"example.com/ticket/utilities"
	"github.com/gin-gonic/gin"
)

type SMSRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Message     string `json:"message" binding:"required"`
} 

func SendSmsHandler(context *gin.Context) {
	var req SMSRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"error": "phone_number and message are required"})
		return
	}

	err := utilities.SendSMS(req.PhoneNumber, req.Message)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send SMS", "details": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"status": "SMS sent successfully!"})
}