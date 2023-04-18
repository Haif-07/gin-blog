package front

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Article struct{}

// 根据页数，标签id，分类id
func (*Article) ArticlePageNum(c *gin.Context) {
	var pageinfo models.PageInfo
	pageNum := c.Query("pageNum")
	tagIds := c.Query("tagIds")
	categoryIds := c.Query("categoryIds")
	p, err := strconv.Atoi(pageNum)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	list, total, err := dao.GetFrontArticleListByPage(p, tagIds, categoryIds)
	if err != nil {
		zap.L().Error("查询文章出错", zap.Error(err))
		response.FailWithMessage("查询出错", c)
	}
	pageinfo.Total = total
	pageinfo.PageNum = p
	pageinfo.PageSize = 10
	pageinfo.Data = list
	response.OkWithDetailed(pageinfo, "查询成功", c)
}

// 归档请求
func (*Article) GetArticlesAll(c *gin.Context) {
	all, err := dao.GetArticlesAll()
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		response.FailWithMessage("查询出错", c)
	}
	response.OkWithDetailed(all, "查询成功", c)

}

// 根据id查文章细节请求

func (*Article) GetArticleDetailsById(c *gin.Context) {

	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	articledetails, err := dao.GetArticleDetails(i)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		response.FailWithMessage("查询出错", c)
	}
	previous, err := dao.GetArticlePrevious(i)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
	}
	next, err := dao.GetArticleNext(i)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
	}
	articledetails.Previous = previous
	articledetails.Next = next
	response.OkWithDetailed(articledetails, "查询成功", c)

}

func (*Article) GetArticleComments(c *gin.Context) {
	var pageinfo models.PageInfo
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	pageNum := c.Query("pageNum")
	p, err := strconv.Atoi(pageNum)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	totle, commentslist, err := dao.GetArticleCommentsById(i, p)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		response.FailWithMessage("查询出错", c)
	}
	pageinfo.Total = totle
	pageinfo.PageNum = p
	pageinfo.PageSize = 10
	pageinfo.Data = commentslist
	response.OkWithDetailed(pageinfo, "查询成功", c)

}

func (*Article) SearchArticle(c *gin.Context) {
	var pageNum, pageSize int
	var pageinfo models.PageInfo
	p := c.Query("pageNum")
	if p == "" {
		pageNum = 1
		pageSize = 5
	} else {
		pageSize = 10
		pn, err := strconv.Atoi(p)
		if err != nil {
			zap.L().Error("请求参数转换出错", zap.Error(err))
		}
		pageNum = pn
	}
	queryString := c.Query("queryString")

	list, i, err := dao.Search(pageNum, pageSize, queryString)
	if err != nil {
		zap.L().Error("查询出错", zap.Error(err))
		response.FailWithMessage("查询出错", c)
	}
	pageinfo.Total = i
	pageinfo.PageNum = pageNum
	pageinfo.PageSize = pageSize
	pageinfo.Data = list
	response.OkWithDetailed(pageinfo, "查询成功", c)

}
