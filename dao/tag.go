package dao

import (
	"gin-blog/database"
	"gin-blog/models"
	"gin-blog/utils/mytime"
	"time"
)

func GetTagsWithArticleCount() ([]models.Tagdto, error) {
	//困难，"only_full_group_by"sql模式对GROUP BY的影响及处理
	var taglist []models.Tagdto
	err := database.DB.Table("tags t").
		Select("ANY_VALUE(t.id) as id", "t.name", "COUNT(at.article_id) AS article_count", "t.created_at", "t.updated_at").
		Joins("LEFT JOIN article_tags at ON t.id = at.tag_id").Group("t.name").Find(&taglist).Error
	if err != nil {
		return taglist, err
	}
	return taglist, nil
}

func GetTagCount() (int64, error) {
	var count int64
	if err := database.DB.Model(&models.Tag{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

// 后台删除一条标签
func DeleteTagById(id int) error {
	var tag models.Tag
	tag.Id = id
	if err := database.DB.Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}

// 后台新增一条标签
func AddTag(tag *models.Tag) error {
	tag.CreatedAt = mytime.MyTime(time.Now())
	tag.UpdatedAt = mytime.MyTime(time.Now())

	if err := database.DB.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}
