package dao

import (
	"gin-blog/database"
	"gin-blog/models"
	"time"
)

func GetCategoryWithArticleCount() []models.Categorydto {
	//困难，"only_full_group_by"sql模式对GROUP BY的影响及处理
	var clist []models.Categorydto
	database.DB.Table("categories c").
		Select("ANY_VALUE(c.id) as id", "c.name", "COUNT(ac.article_id) AS article_count", "c.created_at", "c.updated_at").
		Joins("LEFT JOIN article_categories ac ON c.id = ac.category_id").Group("c.name").Find(&clist)
	return clist
}

func GetCategoryCount() int {
	var count int64
	database.DB.Model(&models.Category{}).Count(&count)
	return int(count)
}

// 后台新增一条分类
func AddCategory(category *models.Category) int {
	category.CreatedAt = models.MyTime(time.Now())
	category.UpdatedAt = models.MyTime(time.Now())
	row := database.DB.Create(&category)
	i := row.RowsAffected
	return int(i)

}

// 后台删除一台分类
func DeleteCategoryById(id int) int64 {

	var cat models.Category
	cat.Id = id

	row := database.DB.Delete(&cat)
	i := row.RowsAffected
	return i

}
