package controllers

import (
	"github.com/MrWildanMD/room-booking/models"
	"github.com/gin-gonic/gin"
)

func sendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"status": "error", "message": message})
}

func sendSuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	response := gin.H{"status": "success", "message": message}
	if data != nil {
		response["data"] = data
	}
	c.JSON(status, response)
}

func IsAdmin(user models.Users) bool {
	return user.TypeUser == 1
}

func IsApproved(booking models.Booking) bool {
	return booking.BookingStatus == 1
}