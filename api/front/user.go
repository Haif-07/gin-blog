package front

import (
	"gin-blog/database"
	"gin-blog/models"
	"gin-blog/models/response"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) GetUserInfo(c *gin.Context) {

}

func (*User) UserInfo(c *gin.Context) {
	var userinfo models.User
	userSocialUserId, flag := c.Get("user")
	if !flag {
		zap.L().Error("权限不足，没有存储到信息")
		response.FailWithMessage("查询出错了", c)
	}
	err := database.DB.Where("social_user_id = ?", userSocialUserId).First(&userinfo).Error
	if err != nil {
		zap.L().Error("查询出错了", zap.Error(err))
		response.FailWithMessage("查询出错了", c)
	}
	response.OkWithDetailed(userinfo, "查询成功", c)
}
func (*User) FrontLogin(c *gin.Context) {

}
