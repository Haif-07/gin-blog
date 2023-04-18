package models

import (
	"gin-blog/utils/mytime"
)

type Category struct {
	Id                        int                         `json:"id"`
	Name                      string                      `json:"name"`
	CreatedAt                 mytime.MyTime               `json:"createdAt"`
	UpdatedAt                 mytime.MyTime               `json:"updatedAt"`
	CreatedOrUpdateArticleDto []CreatedOrUpdateArticleDto `gorm:"many2many:article_categories"`
}

type Categorydto struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	ArticleCount int           `json:"articleCount"`
	CreatedAt    mytime.MyTime `json:"createdAt"`
	UpdatedAt    mytime.MyTime `json:"updatedAt"`
}
