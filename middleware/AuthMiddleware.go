package middleware

import (
	"gin-blog/models/response"
	"gin-blog/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenstring := ctx.GetHeader("Authorization")

		if tokenstring == "" || !strings.HasPrefix(tokenstring, "Bearer ") {
			response.FailWithMessage("权限不足", ctx)
			ctx.Abort()
			return
		}

		tokenstring = tokenstring[7:]
		token, cliams, err := utils.ParseToken(tokenstring)

		if err != nil || !token.Valid {
			response.FailWithMessage("权限不足", ctx)
			ctx.Abort()
			return
		}
		userSocialUserId := cliams.SocialUserId
		ctx.Set("user", userSocialUserId)
		ctx.Next()
	}
}
