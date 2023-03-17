package front

import (
	"gin-blog/dao"
	"gin-blog/models"
	"net/http"
	"strconv"

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
		return
	}
	list, total := dao.GetFrontArticleListByPage(p, tagIds, categoryIds)
	pageinfo.Total = total
	pageinfo.PageNum = p
	pageinfo.PageSize = 10
	pageinfo.Data = list
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": pageinfo,
	})

}

// 归档请求
func (*Article) GetArticlesAll(c *gin.Context) {
	list := dao.GetArticlesAll()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": list,
	})
}

// 根据id查文章细节请求

func (*Article) GetArticleDetailsById(c *gin.Context) {

	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	articledetails := dao.GetArticleDetails(i)
	articledetails.Previous = dao.GetArticlePrevious(i)
	articledetails.Next = dao.GetArticleNext(i)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": articledetails,
	})
}

func (*Article) GetArticleComments(c *gin.Context) {
	var pageinfo models.PageInfo
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}
	pageNum := c.Query("pageNum")
	p, err := strconv.Atoi(pageNum)
	if err != nil {
		return
	}
	totle, commentslist := dao.GetArticleCommentsById(i, p)
	pageinfo.Total = totle
	pageinfo.PageNum = p
	pageinfo.PageSize = 10
	pageinfo.Data = commentslist
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": pageinfo,
	})
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
			return
		}
		pageNum = pn

	}

	queryString := c.Query("queryString")

	list, i := dao.Search(pageNum, pageSize, queryString)

	pageinfo.Total = i
	pageinfo.PageNum = pageNum
	pageinfo.PageSize = pageSize
	pageinfo.Data = list
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": pageinfo,
	})

}
