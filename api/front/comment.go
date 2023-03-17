package front

import (
	"gin-blog/dao"
	"gin-blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Comment struct{}

func (*Comment) Save(c *gin.Context) {

	var commentreq models.CommentReq
	err := c.ShouldBindJSON(&commentreq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i := dao.SaveComment(commentreq)
	if i > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "fail",
		})
	}

}

func (*Comment) DelectComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	i := dao.DelectCommentById(id)
	if i > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "fail",
		})
	}

}
