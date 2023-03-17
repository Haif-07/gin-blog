package front

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) GetUserInfo(c *gin.Context) {

}

func (*User) UserInfo(c *gin.Context) {
	userinfo, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"data": userinfo,
	})
}
func (*User) FrontLogin(c *gin.Context) {
	// var u models.User
	// c.ShouldBindJSON(&u)
	// fmt.Println(u)
	// code := dao.FrontLogin(&u)
	// if code == errmsg.SUCCSE {

	// 	token, err := utils.GetToken(u)
	// 	if err != nil {
	// 		return
	// 	}

	// 	c.JSON(200, gin.H{
	// 		"data": gin.H{
	// 			"Oauth-Token": token,
	// 		},
	// 		"status":  code,
	// 		"message": errmsg.GetErrmsg(code),
	// 	})
	// } else {
	// 	c.JSON(500, gin.H{
	// 		"status":  code,
	// 		"message": errmsg.GetErrmsg(code),
	// 	})
	// }

}
