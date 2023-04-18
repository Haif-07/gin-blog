package v1

import (
	"gin-blog/dao"
	"gin-blog/models"
	"gin-blog/models/response"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Tag struct{}

func (*Tag) GetTagALL(c *gin.Context) {
	list, err := dao.GetTagsWithArticleCount()
	if err != nil {
		zap.L().Error("查询出错了", zap.Error(err))
		response.FailWithMessage("查询出错了", c)
	}
	response.OkWithDetailed(list, "查询成功", c)
}

func (*Tag) AddTag(c *gin.Context) {
	var tag models.Tag
	err := c.ShouldBind(&tag)
	if err != nil {
		zap.L().Error("请求参数绑定出错", zap.Error(err))
	}
	err = dao.AddTag(&tag)
	if err != nil {
		zap.L().Error("创建出错了", zap.Error(err))
		response.FailWithMessage("创建出错了", c)
	}
	response.OkWithMessage("创建成功", c)

}
func (*Tag) DeleteTag(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		zap.L().Error("请求参数转换出错", zap.Error(err))
	}
	err = dao.DeleteTagById(id)
	if err != nil {
		zap.L().Error("删除分类出错了", zap.Error(err))
		response.FailWithMessage("删除出错了", c)
	}
	response.OkWithMessage("删除成功", c)
}
