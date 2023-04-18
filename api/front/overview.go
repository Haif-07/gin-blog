package front

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Overview struct{}
type ResourceCount struct {
	Article  int `json:"article"`
	Comment  int `json:"comment"`
	Category int `json:"category"`
	Tag      int `json:"tag"`
}

type OverviewVo struct {
	ResourceCount ResourceCount        `json:"resourceCount"`
	Tags          []models.Tagdto      `json:"tags"`
	Categories    []models.Categorydto `json:"categories"`
	Pages         []models.Page        `json:"pages"`
	Configs       []models.Config      `json:"configs"`
}

func (*Overview) Overview(c *gin.Context) {
	var o OverviewVo
	articlecount, err := dao.GetArticleCount()
	if err != nil {
		zap.L().Error("获取文章总数出错", zap.Error(err))
	}
	categoryCount, err := dao.GetCategoryCount()
	if err != nil {
		zap.L().Error("获取分类总数出错", zap.Error(err))
	}
	commentcount, err := dao.GetCommentCount()
	if err != nil {
		zap.L().Error("获取评论总数出错", zap.Error(err))
	}
	tagcount, err := dao.GetTagCount()
	if err != nil {
		zap.L().Error("获取标签总数出错", zap.Error(err))
	}
	taglist, err := dao.GetTagsWithArticleCount()
	if err != nil {
		zap.L().Error("获取标签出错了", zap.Error(err))
	}
	categorylist, err := dao.GetCategoryWithArticleCount()
	if err != nil {
		zap.L().Error("获取分类出错了", zap.Error(err))
	}
	page, err := dao.GetPage()
	if err != nil {
		zap.L().Error("获取页面信息出错了", zap.Error(err))
	}
	configs, err := dao.Getconfigs()
	if err != nil {
		zap.L().Error("获取博客配置出错了", zap.Error(err))
	}
	o.ResourceCount.Article = int(articlecount)
	o.ResourceCount.Category = int(categoryCount)
	o.ResourceCount.Comment = int(commentcount)
	o.ResourceCount.Tag = int(tagcount)
	o.Tags = taglist
	o.Categories = categorylist
	o.Pages = page
	o.Configs = configs
	response.OkWithDetailed(o, "查询成功", c)

}
