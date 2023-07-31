package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/panbhatt/go-gin-crud-gorm/initializers"
	"github.com/panbhatt/go-gin-crud-gorm/models"
	"github.com/panbhatt/go-gin-crud-gorm/utils"
)

func DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authzHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authzHeader)

		if len(fields) > 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		config := initializers.CO
		sub, err := utils.VerifyToken(access_token, config.AccessTokenPublicKey)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		var user models.User
		result := initializers.DB.First(&user, "id =?", fmt.Sprint(sub))

		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the User does not exists"})
			return

		}

	}
}
