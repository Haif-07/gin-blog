package dao

import (
	"fmt"
	"gin-blog/database"
	"gin-blog/models"
	"time"
)

func GetTagsWithArticleCount() []models.Tagdto {
	//困难，"only_full_group_by"sql模式对GROUP BY的影响及处理
	var taglist []models.Tagdto
	database.DB.Table("tags t").
		Select("ANY_VALUE(t.id) as id", "t.name", "COUNT(at.article_id) AS article_count", "t.created_at", "t.updated_at").
		Joins("LEFT JOIN article_tags at ON t.id = at.tag_id").Group("t.name").Find(&taglist)
	return taglist
}

func GetTagCount() int {
	var count int64
	database.DB.Model(&models.Tag{}).Count(&count)
	return int(count)
}

// 后台删除一条标签
func DeleteTagById(id int) int64 {

	var tag models.Tag
	tag.Id = id

	row := database.DB.Delete(&tag)
	i := row.RowsAffected
	return i

}

// 后台新增一条标签
func AddTag(tag *models.Tag) int {
	tag.CreatedAt = models.MyTime(time.Now())
	tag.UpdatedAt = models.MyTime(time.Now())
	fmt.Printf("tag: %v\n", tag)
	row := database.DB.Create(&tag)
	i := row.RowsAffected
	return int(i)

}
