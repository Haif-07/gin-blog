package dao

import (
	"gin-blog/database"
	"gin-blog/models"
)

func Getconfigs() []models.Config {
	var con []models.Config
	database.DB.Find(&con)
	return con
}
