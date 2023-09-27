package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
    "github.com/MrWildanMD/room-booking/models"
)

type NotificationController struct {
    DB *gorm.DB
}

func (nc *NotificationController) AddNotification(c *gin.Context) {
    var notification models.Notification

    if err := c.ShouldBindJSON(&notification); err != nil {
        sendErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := nc.DB.Create(&notification).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    sendSuccessResponse(c, http.StatusCreated, "Notification added successfully", notification)
}

func (nc *NotificationController) GetNotification(c *gin.Context) {
    var notifications []models.Notification
    if err := nc.DB.Find(&notifications).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    sendSuccessResponse(c, http.StatusOK, "Notifications fetched successfully", notifications)
}

func (nc *NotificationController) GetNotificationByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
        return
    }

    var notification models.Notification
    if err := nc.DB.First(&notification, id).Error; err != nil {
        sendErrorResponse(c, http.StatusNotFound, "Notification not found")
        return
    }

    sendSuccessResponse(c, http.StatusOK, "Notification fetched successfully", notification)
}

func (nc *NotificationController) UpdateNotification(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
        return
    }

    var notification models.Notification
    if err := nc.DB.First(&notification, id).Error; err != nil {
        sendErrorResponse(c, http.StatusNotFound, "Notification not found")
        return
    }

    if err := c.ShouldBindJSON(&notification); err != nil {
        sendErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := nc.DB.Save(&notification).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    sendSuccessResponse(c, http.StatusOK, "Notification updated successfully", notification)
}

func (nc *NotificationController) DeleteNotification(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
        return
    }

    var notification models.Notification
    if err := nc.DB.First(&notification, id).Error; err != nil {
        sendErrorResponse(c, http.StatusNotFound, "Notification not found")
        return
    }

    if err := nc.DB.Delete(&notification).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    sendSuccessResponse(c, http.StatusNoContent, "Notification deleted successfully", nil)
}