package dao

import (
	"gin-blog/database"
	"gin-blog/models"
	"gin-blog/utils/mytime"
	"time"
)

func GetCategoryWithArticleCount() ([]models.Categorydto, error) {
	//困难，"only_full_group_by"sql模式对GROUP BY的影响及处理
	var clist []models.Categorydto
	err := database.DB.Table("categories c").
		Select("ANY_VALUE(c.id) as id", "c.name", "COUNT(ac.article_id) AS article_count", "c.created_at", "c.updated_at").
		Joins("LEFT JOIN article_categories ac ON c.id = ac.category_id").
		Group("c.name").
		Find(&clist).Error
	if err != nil {
		return clist, err
	}
	return clist, nil
}

func GetCategoryCount() (int64, error) {
	var count int64
	if err := database.DB.Model(&models.Category{}).Count(&count).Error; err != nil {
		return count, err
	}
	return count, nil
}

// 后台新增一条分类
func AddCategory(category *models.Category) error {
	category.CreatedAt = mytime.MyTime(time.Now())
	category.UpdatedAt = mytime.MyTime(time.Now())
	if err := database.DB.Create(&category).Error; err != nil {
		return err
	}
	return nil

}

// 后台删除一台分类
func DeleteCategoryById(id int) error {

	var cat models.Category
	cat.Id = id

	if err := database.DB.Delete(&cat).Error; err != nil {
		return err
	}
	return nil

}
