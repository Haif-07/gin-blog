package models

import (
	"gin-blog/utils/mytime"
)

type Article struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Content     string        `json:"content"`
	Status      int           `json:"status"`
	CreatedAt   mytime.MyTime `json:"createdAt"`
	UpdatedAt   mytime.MyTime `json:"updatedAt"`
}

type CreatedOrUpdateArticleDto struct {
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Content     string        `json:"content"`
	Status      int           `json:"status"`
	CreatedAt   mytime.MyTime `json:"createdAt"`
	UpdatedAt   mytime.MyTime `json:"updatedAt"`
	Tags        []Tag         `json:"tags,omitempty" gorm:"many2many:article_tags;joinForeignKey:article_id;"`
	Categories  []Category    `json:"categories,omitempty" gorm:"many2many:article_categories;joinForeignKey:article_id;"`
}

// 首页响应的结构体
type ArticleCategoryAndTag struct {
	ID          int           `json:"id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Status      int           `json:"status,omitempty"`
	CreatedAt   mytime.MyTime `json:"createdAt,omitempty"`
	UpdatedAt   mytime.MyTime `json:"updatedAt,omitempty"`
	Tags        []Tag         `json:"tags,omitempty" gorm:"many2many:article_tags;joinForeignKey:article_id;"`
	Categories  []Category    `json:"categories,omitempty" gorm:"many2many:article_categories;joinForeignKey:article_id;"`
}

// 归档的结构体
type ArticleSomeInfo struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	CreatedAt mytime.MyTime `json:"createdAt"`
}

// 文章详细页的响应结构体
type ArticleDetails struct {
	ID          int               `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Content     string            `json:"content"`
	Status      int               `json:"status"`
	CreatedAt   mytime.MyTime     `json:"createdAt"`
	UpdatedAt   mytime.MyTime     `json:"updatedAt"`
	Tags        []Tag             `json:"tags" gorm:"many2many:article_tags;joinForeignKey:article_id;"`
	Categories  []Category        `json:"categories" gorm:"many2many:article_categories;joinForeignKey:article_id;"`
	Previous    ArticlePagination `gorm:"-" json:"previous"`
	Next        ArticlePagination `gorm:"-" json:"next"`
}

type ArticlePagination struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
