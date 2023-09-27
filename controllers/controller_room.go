package controllers

import (
	"net/http"
	"strconv"

	"github.com/MrWildanMD/room-booking/models"
	"github.com/MrWildanMD/room-booking/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomController struct {
	DB *gorm.DB
}

func (rc *RoomController) AddRoom(c *gin.Context) {
	authorizedUser := c.MustGet("authorizedUser").(models.Users)

	if utils.IsEmptyString(authorizedUser.Name) {
		sendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	if !IsAdmin(authorizedUser) {
		sendErrorResponse(c, http.StatusForbidden, "Permission Denied")
		return
	}

	var room models.Room

	if err := c.ShouldBindJSON(&room); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := rc.DB.Create(&room).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusCreated, "Room added successfully", room)
}

func (rc *RoomController) GetRoom(c *gin.Context) {
	var rooms []models.Room
	if err := rc.DB.Find(&rooms).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusOK, "Rooms fetched successfully", rooms)
}

func (rc *RoomController) GetRoomByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var room models.Room
	if err := rc.DB.First(&room, id).Error; err != nil {
		sendErrorResponse(c, http.StatusNotFound, "Room not found")
		return
	}

	sendSuccessResponse(c, http.StatusOK, "Room fetched successfully", room)
}

func (rc *RoomController) UpdateRoom(c *gin.Context) {
	authorizedUser := c.MustGet("authorizedUser").(models.Users)

	if utils.IsEmptyString(authorizedUser.Name) {
		sendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	if !IsAdmin(authorizedUser) {
		sendErrorResponse(c, http.StatusForbidden, "Permission Denied")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var room models.Room
	if err := rc.DB.First(&room, id).Error; err != nil {
		sendErrorResponse(c, http.StatusNotFound, "Room not found")
		return
	}

	if err := c.ShouldBindJSON(&room); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := rc.DB.Save(&room).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusOK, "Room updated successfully", room)
}

func (rc *RoomController) DeleteRoom(c *gin.Context) {
	authorizedUser := c.MustGet("authorizedUser").(models.Users)

	if utils.IsEmptyString(authorizedUser.Name) {
		sendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

	if !IsAdmin(authorizedUser) {
		sendErrorResponse(c, http.StatusForbidden, "Permission Denied")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var room models.Room
	if err := rc.DB.First(&room, id).Error; err != nil {
		sendErrorResponse(c, http.StatusNotFound, "Room not found")
		return
	}

	if err := rc.DB.Delete(&room).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusNoContent, "Room deleted successfully", nil)
}
