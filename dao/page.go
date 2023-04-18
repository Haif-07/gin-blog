package dao

import (
	"gin-blog/database"
	"gin-blog/models"
)

func GetPage() ([]models.Page, error) {

	var p []models.Page
	if err := database.DB.Find(&p).Error; err != nil {
		return p, err
	}
	return p, nil
}
