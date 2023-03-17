package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/utils"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	u, err := dao.GetUser(params.Username, params.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}
	token, err := utils.GetToken(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500000,
			"message": "fail",
		})
	}
	c.Header("AUTHORIZATION", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": u,
	})
}

// 后台查看所有用户
func (*User) GetUserList(c *gin.Context) {
	var pageinfo models.PageInfo
	pn := c.Query("pageNum")
	ps := c.Query("pageSize")
	pageNum, err := strconv.Atoi(pn)
	if err != nil {
		return
	}
	pageSize, err := strconv.Atoi(ps)
	if err != nil {
		return
	}
	userlist, total := dao.GetUserList(pageNum, pageSize)
	pageinfo.Total = total
	pageinfo.PageNum = pageNum
	pageinfo.PageSize = pageSize
	pageinfo.Data = userlist
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": pageinfo,
	})

}

// 后台查看某个用户详细信息
func (*User) GetUserById(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	userinfo, err := dao.GetUserById(i)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": userinfo,
	})

}
