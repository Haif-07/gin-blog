package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Category struct{}

func (*Category) GetCategoryALL(c *gin.Context) {
	list := dao.GetCategoryWithArticleCount()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": list,
	})
}

func (*Category) AddCategory(c *gin.Context) {
	var category models.Category
	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	i := dao.AddCategory(&category)
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
func (*Category) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	i := dao.DeleteCategoryById(id)
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
