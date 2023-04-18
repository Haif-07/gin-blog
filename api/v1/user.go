package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"
	"gin-blog/utils"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type User struct{}

// 登录参数绑定
type UserParams struct {
	Username string
	Password string
}

// 后台登录
func (*User) AuthLogin(c *gin.Context) {
	var params UserParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		zap.L().Error("请求参数绑定出错", zap.Error(err))
	}
	u, err := dao.GetUser(params.Username, params.Password)
	if err != nil {
		zap.L().Error("查询用户出错了", zap.Error(err))
		response.FailWithMessage("查询出错了", c)
	}
	token, err := utils.GetToken(u)
	if err != nil {
		zap.L().Error("创建token出错", zap.Error(err))
		response.FailWithMessage("fail", c)
	}
	c.Header("AUTHORIZATION", "Bearer "+token)
	response.OkWithDetailed(u, "success", c)

}

// 后台查看所有用户
func (*User) GetUserList(c *gin.Context) {
	var pageinfo models.PageInfo
	pn := c.Query("pageNum")
	ps := c.Query("pageSize")
	pageNum, err := strconv.Atoi(pn)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	userlist, total, err := dao.GetUserList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("查询所有用户出错了", zap.Error(err))
		response.FailWithMessage("查询出错了", c)
	}
	pageinfo.Total = total
	pageinfo.PageNum = pageNum
	pageinfo.PageSize = pageSize
	pageinfo.Data = userlist
	response.OkWithDetailed(pageinfo, "查询成功", c)

}

// 后台查看某个用户详细信息
func (*User) GetUserById(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	userinfo, err := dao.GetUserById(i)
	if err != nil {
		zap.L().Error("查询用户信息出错了", zap.Error(err))
		response.FailWithMessage("查询出错了", c)
	}
	response.OkWithDetailed(userinfo, "查询成功", c)

}
