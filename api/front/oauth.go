package front

import (
	"gin-blog/config"
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"
	"gin-blog/utils"
	"gin-blog/utils/mytime"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Oauth struct{}

func (*Oauth) Login(c *gin.Context) {

	// url := config.AuthCodeURL("state", oauth2.AccessTypeOnline)
	authinfo := config.Setup()

	url := config.GetUrl(authinfo)

	c.JSON(http.StatusOK, gin.H{
		"code":         200000,
		"message":      "success",
		"authorizeUrl": url,
	})

}
func (*Oauth) GithubLogincallback(c *gin.Context) {
	code := c.Query("code")
	authinfo := config.Setup()
	tok, err := config.GetToken(c, authinfo, code)
	if err != nil {
		zap.L().Error("获取token出错", zap.Error(err))
	}
	client := authinfo.Client(c, tok)
	u, err := config.GetUsers(client)
	if err != nil {
		zap.L().Error("获取第三方信息出错", zap.Error(err))
	}

	var userinfo models.User
	userinfo.SocialSource = "github"
	userinfo.SocialUserId = strconv.Itoa(u.ID)
	userinfo.Username = u.Login
	userinfo.AvatarUrl = u.AvatarURL
	userinfo.Email = u.Email
	userinfo.Role = "1"
	userinfo.LastLogin = mytime.MyTime(time.Now())

	err = dao.CreatedOrUpdate(&userinfo)
	if err != nil {
		zap.L().Error("创建或更新用户出错", zap.Error(err))
		response.FailWithMessage("fail", c)
	}
	token, err := utils.GetToken(userinfo)
	if err != nil {
		zap.L().Error("创建token出错", zap.Error(err))
		response.FailWithMessage("fail", c)
	}
	c.Header("AUTHORIZATION", "Bearer "+token)

	response.OkWithMessage("success", c)

}
