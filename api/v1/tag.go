package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func (*Tag) GetTagALL(c *gin.Context) {
	list := dao.GetTagsWithArticleCount()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": list,
	})
}

func (*Tag) AddTag(c *gin.Context) {
	var tag models.Tag
	err := c.ShouldBind(&tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i := dao.AddTag(&tag)
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
func (*Tag) DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	i := dao.DeleteTagById(id)
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
