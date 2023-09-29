package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/MrWildanMD/room-booking/config"
	"github.com/MrWildanMD/room-booking/models"
	"github.com/MrWildanMD/room-booking/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type BookingController struct {
	DB *gorm.DB
}

func (bc *BookingController) AddBooking(c *gin.Context) {
	var booking models.Booking

	if err := c.ShouldBindJSON(&booking); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	booking.BookingStatus = 0

	if err := bc.DB.Create(&booking).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	var user models.Users
	bc.DB.First(&user, "id = ?", booking.UserID)
	var room models.Room
	bc.DB.First(&room, "id = ?", booking.RoomID)

	sendSuccessResponse(c, http.StatusCreated, "Booking added successfully", booking)
	message := "New Booking Added! \nUser: " + user.Name + "\nRoom: " + room.Title + "\nCheck In: " + booking.CheckIn.String() + "\nCheck Out: " + booking.CheckOut.String() + "\nGuest Total: " + utils.IntToString(booking.NumberOfGuests) + "\nApproved: " + utils.BoolToString(IsApproved(booking)) + "\nAdditional Item: " + booking.AdditionalItem + "\nTimestamp: " + time.Now().String()
	config.TB.Send(&telebot.User{ID: config.ChatID}, message)
}

func (bc *BookingController) GetBooking(c *gin.Context) {
	var bookings []models.Booking
	if err := bc.DB.Preload("User").Preload("Room").Find(&bookings).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

	sendSuccessResponse(c, http.StatusOK, "Bookings fetched successfully", bookings)
}

func (bc *BookingController) GetBookingByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var booking models.Booking
	if err := bc.DB.Preload("User").Preload("Room").First(&booking, id).Error; err != nil {
        sendErrorResponse(c, http.StatusNotFound, "Booking not found")
        return
    }

	sendSuccessResponse(c, http.StatusOK, "Booking fetched successfully", booking)
}

func (bc *BookingController) UpdateBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var booking models.Booking
	if err := bc.DB.Preload("User").Preload("Room").First(&booking, id).Error; err != nil {
        sendErrorResponse(c, http.StatusNotFound, "Booking not found")
        return
    }

	if err := c.ShouldBindJSON(&booking); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := bc.DB.Save(&booking).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var user models.Users
	bc.DB.First(&user, "id = ?", booking.UserID)
	var room models.Room
	bc.DB.First(&room, "id = ?", booking.RoomID)

	sendSuccessResponse(c, http.StatusOK, "Booking updated successfully", booking)
	message := "Booking Updated! \nUser: " + user.Name + "\nRoom: " + room.Title + "\nCheck In: " + booking.CheckIn.String() + "\nCheck Out: " + booking.CheckOut.String() + "\nGuest Total: " + utils.IntToString(booking.NumberOfGuests) + "\nApproved: " + utils.BoolToString(IsApproved(booking)) + "\nAdditional Item: " + booking.AdditionalItem + "\nTimestamp: " + time.Now().String()
	config.TB.Send(&telebot.User{ID: config.ChatID}, message)
}

func (bc *BookingController) DeleteBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var booking models.Booking
	if err := bc.DB.First(&booking, id).Error; err != nil {
		sendErrorResponse(c, http.StatusNotFound, "Booking not found")
		return
	}

	var user models.Users
	bc.DB.First(&user, "id = ?", booking.UserID)
	var room models.Room
	bc.DB.First(&room, "id = ?", booking.RoomID)

	if err := bc.DB.Delete(&booking).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusNoContent, "Booking deleted successfully", nil)
	message := "Booking Deleted! \nUser: " + user.Name + "\nRoom: " + room.Title + "\nCheck In: " + booking.CheckIn.String() + "\nCheck Out: " + booking.CheckOut.String() + "\nGuest Total: " + utils.IntToString(booking.NumberOfGuests) + "\nApproved: " + utils.BoolToString(IsApproved(booking)) + "\nAdditional Item: " + booking.AdditionalItem + "\nTimestamp: " + time.Now().String()
	config.TB.Send(&telebot.User{ID: config.ChatID}, message)
}
