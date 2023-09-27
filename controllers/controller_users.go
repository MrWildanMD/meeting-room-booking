package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MrWildanMD/room-booking/config"
	"github.com/MrWildanMD/room-booking/models"
	"github.com/MrWildanMD/room-booking/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersController struct {
	DB *gorm.DB
}

type LoginRequest struct {
	PrivyID string `json:"privy_id"`
	Email   string `json:"email"`
}

type RegisterRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	PrivyID string `json:"privy_id"`
}

func (uc *UsersController) RegisterUsers(c *gin.Context) {
	var body *RegisterRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.Users
	user.Name = body.Name
	user.Email = body.Email
	user.PrivyID = body.PrivyID
	user.TypeUser = 0

	res := uc.DB.Create(&user)
	if res.Error != nil {
		sendErrorResponse(c, http.StatusBadGateway, "Something went wrong!")
		return
	}

	sendSuccessResponse(c, http.StatusCreated, "Registered Successfully, You can login with the credentials registered", nil)
}

func (uc *UsersController) LoginUsers(c *gin.Context) {
	var body *LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var user models.Users
	res := uc.DB.First(&user, "email = ? and privy_id = ?", body.Email, body.PrivyID)
	if res.Error != nil {
		sendErrorResponse(c, http.StatusBadRequest, "Invalid email or privy id")
		return
	}
	cfg, _ := config.LoadConfig()
	token, err := utils.GenerateToken(user.ID, cfg.TOKEN_SECRET)
	if err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.SetCookie("token", token, cfg.TOKEN_MAX_AGE*60, "/", "localhost", false, true)
	sendSuccessResponse(c, http.StatusOK, "Login Success", token)
}

func (uc *UsersController) LogoutUsers(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			sendErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("%v", r))
		}
	}()
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	sendSuccessResponse(c, http.StatusOK, "Logout Successfully", nil)
}

func (uc *UsersController) GetUsers(c *gin.Context) {
    var users []models.Users
    if err := uc.DB.Find(&users).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    sendSuccessResponse(c, http.StatusOK, "Users fetched successfully", users)
}

func (uc *UsersController) GetMe(c *gin.Context) {
    authorizedUser := c.MustGet("authorizedUser").(models.Users)

    if utils.IsEmptyString(authorizedUser.Name) {
		sendErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	}

    // var user models.Users
    // if err := uc.DB.First(&user, id).Error; err != nil {
    //     sendErrorResponse(c, http.StatusNotFound, "User not found")
    //     return
    // }

    sendSuccessResponse(c, http.StatusOK, "User fetched successfully", authorizedUser)
}

func (uc *UsersController) UpdateUsers(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        sendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
        return
    }

    var user models.Users
    if err := uc.DB.First(&user, id).Error; err != nil {
        sendErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }

    var updateUser models.Users
    if err := c.ShouldBindJSON(&updateUser); err != nil {
        sendErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }
	
	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.PrivyID = updateUser.PrivyID

    if err := uc.DB.Save(&user).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    sendSuccessResponse(c, http.StatusOK, "User updated successfully", user)
}

func (uc *UsersController) DeleteUsers(c *gin.Context) {
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

    var user models.Users
    if err := uc.DB.First(&user, id).Error; err != nil {
        sendErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }

    if err := uc.DB.Delete(&user).Error; err != nil {
        sendErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    sendSuccessResponse(c, http.StatusNoContent, "User deleted successfully", nil)
}

