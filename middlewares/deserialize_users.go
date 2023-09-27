package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/MrWildanMD/room-booking/config"
	"github.com/MrWildanMD/room-booking/database"
	"github.com/MrWildanMD/room-booking/models"
	"github.com/MrWildanMD/room-booking/utils"
	"github.com/gin-gonic/gin"
)

func DeserializeUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		cookie, err := c.Cookie("token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		} else if err != nil {
			token = cookie
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}
		config, _ := config.LoadConfig()
		sub, err := utils.ValidateToken(token, config.TOKEN_SECRET)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.Users
		result := database.DB.First(&user, "id = ?", fmt.Sprint(sub))
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "user token expired"})
			return
		}

		c.Set("authorizedUser", user)
		c.Next()
	}
}