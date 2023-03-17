package models

type Category struct {
	Id                        int                         `json:"id"`
	Name                      string                      `json:"name"`
	CreatedAt                 MyTime                      `json:"createdAt"`
	UpdatedAt                 MyTime                      `json:"updatedAt"`
	CreatedOrUpdateArticleDto []CreatedOrUpdateArticleDto `gorm:"many2many:article_categories"`
}

type Categorydto struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ArticleCount int    `json:"articleCount"`
	CreatedAt    MyTime `json:"createdAt"`
	UpdatedAt    MyTime `json:"updatedAt"`
}
