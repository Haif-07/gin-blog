package dao

import (
	"gin-blog/database"
	"gin-blog/models"
)

func GetPage() []models.Page {
	//困难，"only_full_group_by"sql模式对GROUP BY的影响及处理
	var p []models.Page
	database.DB.Find(&p)
	return p
}
