package dao

import (
	"fmt"
	"gin-blog/database"
	"gin-blog/models"

	"gorm.io/gorm"
)

func GetArticleCount() int {
	var count int64
	database.DB.Model(&models.Article{}).Count(&count)
	return int(count)
}

// 前台首页根据页码查询 or 标签id、分类id
func GetFrontArticleListByPage(pageNum int, tagIds, categoryIds string) ([]models.ArticleCategoryAndTag, int64) {
	var total int64
	list := make([]models.ArticleCategoryAndTag, 0)
	db := database.DB.Table("articles")
	if tagIds != "" {
		db = db.Where("id IN (SELECT article_id FROM article_tags WHERE tag_id = ?) ", tagIds)

		fmt.Println("1")
	}
	if categoryIds != "" {
		db = db.Where("id IN (SELECT article_id FROM article_categories WHERE category_id = ?) ", categoryIds)

		fmt.Println("2")
	}
	db.Count(&total)
	db.Preload("Tags").
		Preload("Categories").
		Order("created_at DESC").
		Limit(10).
		Offset(10 * (pageNum - 1)).
		Find(&list)

	return list, total
}

// 根据文章ID获取文章详细信息
func GetArticleDetails(id int) (res models.ArticleDetails) {
	database.DB.Table("articles").
		Preload("Categories").
		Preload("Tags").
		Where("id = ?", id).
		First(&res)
	return
}

// 获取上一个文章
func GetArticlePrevious(id int) (res models.ArticlePagination) {
	sub := database.DB.Table("articles").
		Select("max(id)").
		Where("id < ?", id)
	database.DB.Table("articles").
		Where("id = (?)", sub).
		Limit(1).Find(&res)
	return
}

// 获取下一个文章
func GetArticleNext(id int) (res models.ArticlePagination) {

	database.DB.Table("articles").
		Where("id > ?", id).
		Limit(1).
		Find(&res)
	return
}

// 前台归档
func GetArticlesAll() []models.ArticleSomeInfo {
	list := make([]models.ArticleSomeInfo, 0)
	database.DB.Table("articles").Order("created_at DESC").Find(&list)
	return list
}

// 文章评论的展示
// 比较难，多理解，待优化（可能）
func GetArticleCommentsById(i, p int) (totle int64, commentslist []models.CommentVo) {
	database.DB.
		Table("comments").
		Where("article_id = ? and parent_id  = 0 ", i). //parent_id为null则代表是父评论
		Count(&totle).
		Preload("User").
		Preload("ReplyToUser").
		Preload("SubComments", func(db *gorm.DB) *gorm.DB { //这里subComments的外键是parent_id！！！
			return db.Order("created_at ASC")
		}).
		Preload("SubComments.User").
		Preload("SubComments.ReplyToUser").
		Order("created_at DESC").
		Limit(10).
		Offset(10 * (p - 1)).
		Find(&commentslist)
	return totle, commentslist
}

// 删除
func DeleteArticleById(id int) int64 {
	var art models.Article
	art.Id = id
	row := database.DB.Delete(&art)
	i := row.RowsAffected
	return i
}

//新增博客

func CreateArticle(article *models.CreatedOrUpdateArticleDto) int64 {

	tags := article.Tags
	categories := article.Categories
	for i, t := range tags {
		database.DB.Table("tags").Where("name = ?", t.Name).First(&t)
		tags[i].Id = t.Id
	}
	for i, c := range categories {
		database.DB.Table("categories").Where("name = ?", c.Name).First(&c)
		categories[i].Id = c.Id
	}

	row := database.DB.Debug().Table("articles").Create(&article)

	i := row.RowsAffected
	return i
}

func UpdateArticleById(article *models.CreatedOrUpdateArticleDto) int64 {

	tags := article.Tags
	categories := article.Categories

	for i, t := range tags {

		database.DB.Table("tags").Where("name = ?", t.Name).Find(&t)
		tags[i].Id = t.Id
		if t.Id == 0 {
			database.DB.Create(&t)
			database.DB.Table("tags").Where("name = ?", t.Name).Find(&t)
			tags[i].Id = t.Id

		}
	}
	for i, c := range categories {

		database.DB.Table("categories").Where("name = ?", c.Name).Find(&c)
		categories[i].Id = c.Id

		if c.Id == 0 {
			database.DB.Create(&c)
			database.DB.Table("tags").Where("name = ?", c.Name).Find(&c)
			categories[i].Id = c.Id

		}
	}
	fmt.Printf("tags: %v\n", tags)
	fmt.Printf("categories: %v\n", categories)
	database.DB.Debug().Where("article_id = ?", article.Id).Delete(models.ArticleTag{})
	database.DB.Debug().Where("article_id = ?", article.Id).Delete(models.ArticleCategory{})

	for _, t := range tags {
		database.DB.Create(models.ArticleTag{ArticleId: article.Id, TagId: t.Id})
	}
	for _, c := range categories {
		database.DB.Create(models.ArticleCategory{ArticleId: article.Id, CategoryId: c.Id})
	}

	row := database.DB.Table("articles").Updates(
		models.Article{
			Id: article.Id, Title: article.Title, Description: article.Description,
			Content: article.Content, Status: article.Status, CreatedAt: article.CreatedAt, UpdatedAt: article.UpdatedAt,
		}).Where("id=?", article.Id)

	i := row.RowsAffected
	fmt.Printf("i: %v\n", i)
	return i
}

func Search(pageNum, pageSize int, queryString string) ([]models.ArticleCategoryAndTag, int64) {

	var count int64

	list := make([]models.ArticleCategoryAndTag, 0)

	database.DB.Debug().Table("articles").
		Where(fmt.Sprintf("content like '%%%s'", queryString+"%")).
		Or(fmt.Sprintf("title like '%%%s'", queryString+"%")).
		Or(fmt.Sprintf("description like '%%%s'", queryString+"%")).
		Count(&count).
		Limit(pageSize).
		Offset(10 * (pageNum - 1)).
		Find(&list)
	return list, count
}
