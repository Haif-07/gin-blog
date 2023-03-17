package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Comment struct{}

// 后台管理评论状态统计
func (*Comment) CountCommentNotStatus(c *gin.Context) {
	s := c.Query("status")
	i := dao.CountCommentStatus(s)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": i,
	})

}

func (*Comment) GetCommentsByPage(c *gin.Context) {
	var pageinfo models.PageInfo
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")
	status := c.Query("status")
	if status == "" {
		status = "10"
	}
	pn, err := strconv.Atoi(pageNum)
	if err != nil {
		return
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		return
	}
	st, err := strconv.Atoi(status)
	if err != nil {
		return
	}

	list, i := dao.GetCpmmentsListByPage(pn, ps, st)
	pageinfo.Total = i
	pageinfo.PageNum = pn
	pageinfo.PageSize = ps
	pageinfo.Data = list
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": pageinfo,
	})

}

// 后台修改评论的审核状态
func (*Comment) ChangeStatus(c *gin.Context) {

	i := c.Param("id")
	s := c.Param("status")
	id, err := strconv.Atoi(i)
	if err != nil {
		return
	}
	stauts, err2 := strconv.Atoi(s)
	if err2 != nil {
		return
	}
	row := dao.ChangeStatus(id, stauts)
	if row > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    200000,
			"message": "success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500000,
			"message": "fail",
		})
	}
}
