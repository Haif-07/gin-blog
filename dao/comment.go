package dao

import (
	"gin-blog/database"
	"gin-blog/models"
	"time"
)

func GetCommentCount() int {
	var count int64
	database.DB.Model(&models.Comment{}).Count(&count)
	return int(count)
}

// 发布评论 默认状态是0未审核
func SaveComment(commentreq models.CommentReq) int64 {

	comment := models.Comment{
		UserID:        commentreq.UserId,
		ArticleID:     commentreq.Article.Id,
		Content:       commentreq.Content,
		ReplyToUserID: commentreq.ReplyToUser.Id,
		ParentID:      commentreq.ParentId,
		Status:        0,
		CreatedAt:     models.MyTime(time.Now()),
		UpdatedAt:     models.MyTime(time.Now()),
	}
	row := database.DB.Create(&comment)
	i := row.RowsAffected
	return i
}

// 删除一条评论
func DelectCommentById(id int) int64 {
	var com models.Comment
	com.ID = id

	row := database.DB.Delete(&com)
	i := row.RowsAffected
	return i

}

// 统计评论的不同状态的数量，后台显示的
func CountCommentStatus(s string) int {
	var count int64
	database.DB.Model(&models.Comment{}).Where("status = ?", s).Count(&count)
	return int(count)
}

//后台根据页码获取评论

func GetCpmmentsListByPage(ps, pn, st int) ([]models.Comment, int64) {
	var total int64
	list := make([]models.Comment, 0)
	db := database.DB.Debug().Table("comments")
	if st == 0 {
		db = db.Where("status = ?", st)
	}
	db.Count(&total).
		Preload("User").
		Preload("Article").
		Preload("ReplyToUser").
		Order("created_at DESC").
		Limit(pn).
		Offset(pn * (ps - 1)).
		Find(&list)

	return list, total
}

//后台改评论审核

func ChangeStatus(id, status int) int {

	row := database.DB.Debug().
		Table("comments").Where("id = ?", id).Update("status", status)
	i := row.RowsAffected
	return int(i)
}
