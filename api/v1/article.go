package v1

import (
	"fmt"
	"gin-blog/dao"
	"gin-blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func (*Article) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	i := dao.DeleteArticleById(id)
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

func (*Article) AddArtilce(c *gin.Context) {
	var article models.CreatedOrUpdateArticleDto
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	i := dao.CreateArticle(&article)
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

func (*Article) UpdateArtilce(c *gin.Context) {
	var article models.CreatedOrUpdateArticleDto
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	err2 := c.ShouldBindJSON(&article)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	article.Id = id
	fmt.Printf("article: %v\n", article)
	i := dao.UpdateArticleById(&article)
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
