package front

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Comment struct{}

func (*Comment) Save(c *gin.Context) {
	var commentreq models.CommentReq
	err := c.ShouldBindJSON(&commentreq)
	if err != nil {
		zap.L().Error("请求参数绑定出错", zap.Error(err))
	}
	err = dao.SaveComment(commentreq)
	if err != nil {
		zap.L().Error("保存评论出错了", zap.Error(err))
		response.FailWithMessage("保存评论出错了", c)
	}
	response.OkWithMessage("保存评论成功", c)

}

func (*Comment) DelectComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("请求参数绑定出错", zap.Error(err))
	}
	err = dao.DelectCommentById(id)
	if err != nil {
		zap.L().Error("删除评论出错", zap.Error(err))
		response.FailWithMessage("删除评论出错", c)
	}
	response.OkWithMessage("删除成功", c)

}
