package controllers

import (
	"net/http"
	"strconv"

	"github.com/MrWildanMD/room-booking/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OfficeController struct {
	DB *gorm.DB
}

func (oc *OfficeController) AddOffice(c *gin.Context) {
	var office models.Office

	if err := c.ShouldBindJSON(&office); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := oc.DB.Create(&office).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusCreated, "Office added successfully", office)
}

func (oc *OfficeController) GetOffice(c *gin.Context) {
	var offices []models.Office
	if err := oc.DB.Find(&offices).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusOK, "Offices fetched successfully", offices)
}

func (oc *OfficeController) GetOfficeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var office models.Office
	if err := oc.DB.First(&office, id).Error; err != nil {
		sendErrorResponse(c, http.StatusNotFound, "Office not found")
		return
	}

	sendSuccessResponse(c, http.StatusOK, "Office fetched successfully", office)
}

func (oc *OfficeController) UpdateOffice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var office models.Office
	if err := oc.DB.First(&office, id).Error; err != nil {
		sendErrorResponse(c, http.StatusNotFound, "Office not found")
		return
	}

	if err := c.ShouldBindJSON(&office); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := oc.DB.Save(&office).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusOK, "Office updated successfully", office)
}

func (oc *OfficeController) DeleteOffice(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	var office models.Office
	if err := oc.DB.First(&office, id).Error; err != nil {
		sendErrorResponse(c, http.StatusNotFound, "Office not found")
		return
	}

	if err := oc.DB.Delete(&office).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, http.StatusNoContent, "Office deleted successfully", nil)
}
