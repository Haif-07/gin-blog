package dao

import (
	"gin-blog/database"
	"gin-blog/models"
	"gin-blog/utils/mytime"
	"time"
)

func GetCommentCount() (int64, error) {
	var count int64
	if err := database.DB.Model(&models.Comment{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

// 发布评论 默认状态是0未审核
func SaveComment(commentreq models.CommentReq) error {

	comment := models.Comment{
		UserID:        commentreq.UserId,
		ArticleID:     commentreq.Article.Id,
		Content:       commentreq.Content,
		ReplyToUserID: commentreq.ReplyToUser.Id,
		ParentID:      commentreq.ParentId,
		Status:        0,
		CreatedAt:     mytime.MyTime(time.Now()),
		UpdatedAt:     mytime.MyTime(time.Now()),
	}
	if err := database.DB.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

// 删除一条评论
func DelectCommentById(id int) error {
	var com models.Comment
	com.ID = id
	if err := database.DB.Delete(&com).Error; err != nil {
		return err
	}
	return nil

}

// 统计评论的不同状态的数量，后台显示的
func CountCommentStatus(s string) (int64, error) {
	var count int64
	if err := database.DB.Model(&models.Comment{}).Where("status = ?", s).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

//后台根据页码获取评论

func GetCpmmentsListByPage(ps, pn, st int) ([]models.Comment, int64, error) {
	var total int64
	list := make([]models.Comment, 0)
	db := database.DB.Table("comments")
	if st == 0 {
		db = db.Where("status = ?", st)
	}
	err := db.Count(&total).
		Preload("User").
		Preload("Article").
		Preload("ReplyToUser").
		Order("created_at DESC").
		Limit(pn).
		Offset(pn * (ps - 1)).
		Find(&list).Error
	if err != nil {
		return list, total, err
	}
	return list, total, nil
}

//后台改评论审核

func ChangeStatus(id, status int) error {
	err := database.DB.
		Table("comments").
		Where("id = ?", id).
		Update("status", status).Error
	if err != nil {
		return err
	}
	return nil

}
