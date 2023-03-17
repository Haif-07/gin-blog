package middleware

import (
	"gin-blog/database"
	"gin-blog/models"
	"gin-blog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstring := ctx.GetHeader("Authorization")

		if tokenstring == "" || !strings.HasPrefix(tokenstring, "Bearer ") {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenstring = tokenstring[7:]
		token, cliams, err := utils.ParseToken(tokenstring)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}
		userSocialUserId := cliams.SocialUserId
		var userinfo models.User
		database.DB.Where("social_user_id = ?", userSocialUserId).First(&userinfo)

		ctx.Set("user", userinfo)
		ctx.Next()
	}
}
