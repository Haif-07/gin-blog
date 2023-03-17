package models

type Tag struct {
	Id                        int                         `json:"id"`
	Name                      string                      `json:"name"`
	CreatedAt                 MyTime                      `json:"createdAt"`
	UpdatedAt                 MyTime                      `json:"updatedAt"`
	CreatedOrUpdateArticleDto []CreatedOrUpdateArticleDto `gorm:"many2many:article_tags"`
}
type Tagdto struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ArticleCount int    `json:"articleCount" `
	CreatedAt    MyTime `json:"createdAt"`
	UpdatedAt    MyTime `json:"updatedAt"`
}
