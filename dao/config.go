package dao

import (
	"gin-blog/database"
	"gin-blog/models"
)

func Getconfigs() ([]models.Config, error) {
	var con []models.Config
	if err := database.DB.Find(&con).Error; err != nil {
		return con, err
	}
	return con, nil
}
