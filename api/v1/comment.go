package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Comment struct{}

// 后台管理评论状态统计
func (*Comment) CountCommentNotStatus(c *gin.Context) {
	s := c.Query("status")
	count, err := dao.CountCommentStatus(s)
	if err != nil {
		zap.L().Error("查询未审核评论出错了", zap.Error(err))
		response.FailWithMessage("查询出错", c)
	}

	response.OkWithDetailed(count, "查询成功", c)

}

func (*Comment) GetCommentsByPage(c *gin.Context) {
	var pageinfo models.PageInfo
	pageNum := c.Query("pageNum")
	pageSize := c.Query("pageSize")
	status := c.Query("status")
	if status == "" {
		status = "1"
	}
	pn, err := strconv.Atoi(pageNum)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	st, err := strconv.Atoi(status)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}

	list, count, err := dao.GetCpmmentsListByPage(pn, ps, st)
	if err != nil {
		zap.L().Error("查询评论出错了", zap.Error(err))
		response.FailWithMessage("查询出错了", c)
	}
	pageinfo.Total = count
	pageinfo.PageNum = pn
	pageinfo.PageSize = ps
	pageinfo.Data = list
	response.OkWithDetailed(list, "查询成功", c)

}

// 后台修改评论的审核状态
func (*Comment) ChangeStatus(c *gin.Context) {
	i := c.Param("id")
	s := c.Param("status")
	id, err := strconv.Atoi(i)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	stauts, err := strconv.Atoi(s)
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	err = dao.ChangeStatus(id, stauts)
	if err != nil {
		zap.L().Error("审核评论出错了", zap.Error(err))
		response.FailWithMessage("审核出错了", c)
	}
	response.OkWithMessage("审核成功", c)

}
