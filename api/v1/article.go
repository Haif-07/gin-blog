package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func (*Article) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	err = dao.DeleteArticleById(id)
	if err != nil {
		zap.L().Error("删除用户出错了", zap.Error(err))
		response.FailWithMessage("删除出错了", c)
	}
	response.OkWithMessage("删除成功", c)
}

func (*Article) AddArtilce(c *gin.Context) {
	var article models.CreatedOrUpdateArticleDto
	err := c.ShouldBindJSON(&article)
	if err != nil {
		zap.L().Error("请求参数绑定出错", zap.Error(err))
	}

	err = dao.CreateArticle(&article)
	if err != nil {
		zap.L().Error("创建文章出错了", zap.Error(err))
		response.FailWithMessage("发出出错了", c)
	}
	response.OkWithMessage("发布成功", c)

}

func (*Article) UpdateArtilce(c *gin.Context) {
	var article models.CreatedOrUpdateArticleDto
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	err = c.ShouldBindJSON(&article)
	if err != nil {
		zap.L().Error("请求参数绑定出错", zap.Error(err))
	}
	article.Id = id

	err = dao.UpdateArticleById(&article)
	if err != nil {
		zap.L().Error("更新文章数据出错了", zap.Error(err))
		response.FailWithMessage("编辑出错了", c)
	}
	response.OkWithMessage("编辑成功", c)
}
