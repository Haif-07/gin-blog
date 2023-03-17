package front

import (
	"gin-blog/config"
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/utils"
	"net/http"
	"strconv"
	"time"

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
	// c.Redirect(302, url)
}
func (*Oauth) GithubLogincallback(c *gin.Context) {

	// code := c.Query("code")
	// token, err := config.Exchange(oauth2.NoContext, code)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"access_token": token.AccessToken,
	// })
	code := c.Query("code")
	authinfo := config.Setup()
	tok, err := config.GetToken(c, authinfo, code)
	if err != nil {
		panic(err)
	}
	client := authinfo.Client(c, tok)
	u, err := config.GetUsers(client)
	if err != nil {
		panic(err)
	}

	var userinfo models.User
	userinfo.SocialSource = "github"
	userinfo.SocialUserId = strconv.Itoa(u.ID)
	userinfo.Username = u.Login
	userinfo.AvatarUrl = u.AvatarURL
	userinfo.Email = u.Email
	userinfo.Role = "1"
	userinfo.LastLogin = models.MyTime(time.Now())

	err = dao.CreatedOrUpdate(&userinfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500000,
			"message": "fail",
		})
	}
	// database.DB.Where("social_source=? and social_user_id=?", userinfo.SocialSource, userinfo.SocialUserId).First(&userinfo)
	// if userinfo ==  {
	token, err := utils.GetToken(userinfo)
	// }
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500000,
			"message": "fail",
		})
	}
	c.Header("AUTHORIZATION", "Bearer "+token)

	// fmt.Println(token)
	c.JSON(http.StatusOK, gin.H{
		"code":    200000,
		"message": "success",
	})

}
