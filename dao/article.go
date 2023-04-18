package dao

import (
	"errors"
	"fmt"
	"gin-blog/database"
	"gin-blog/models"

	"gorm.io/gorm"
)

func GetArticleCount() (count int64, err error) {

	err = database.DB.Model(&models.Article{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 前台首页根据页码查询 or 标签id、分类id
func GetFrontArticleListByPage(pageNum int, tagIds, categoryIds string) ([]models.ArticleCategoryAndTag, int64, error) {
	var total int64
	list := make([]models.ArticleCategoryAndTag, 0)
	db := database.DB.Table("articles")
	if tagIds != "" {
		db = db.Where("id IN (SELECT article_id FROM article_tags WHERE tag_id = ?) ", tagIds)
	}
	if categoryIds != "" {
		db = db.Where("id IN (SELECT article_id FROM article_categories WHERE category_id = ?) ", categoryIds)
	}
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Preload("Tags").
		Preload("Categories").
		Order("created_at DESC").
		Limit(10).
		Offset(10 * (pageNum - 1)).
		Find(&list).Error
	if err != nil {
		return list, total, err
	}

	return list, total, nil
}

// 根据文章ID获取文章详细信息
func GetArticleDetails(id int) (models.ArticleDetails, error) {
	var res models.ArticleDetails
	err := database.DB.Table("articles").
		Preload("Categories").
		Preload("Tags").
		Where("id = ?", id).
		First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

// 获取上一个文章
func GetArticlePrevious(id int) (models.ArticlePagination, error) {
	var res models.ArticlePagination
	sub := database.DB.Table("articles").
		Select("max(id)").
		Where("id < ?", id)
	err := database.DB.Table("articles").
		Where("id = (?)", sub).
		Limit(1).Find(&res).Error
	if err != nil {
		return res, errors.New("获取上一篇文章出错")
	}
	return res, nil
}

// 获取下一个文章
func GetArticleNext(id int) (models.ArticlePagination, error) {
	var res models.ArticlePagination
	err := database.DB.Table("articles").
		Where("id > ?", id).
		Limit(1).
		Find(&res).Error
	if err != nil {
		return res, errors.New("获取下一篇文章出错")
	}
	return res, nil
}

// 前台归档
func GetArticlesAll() ([]models.ArticleSomeInfo, error) {
	list := make([]models.ArticleSomeInfo, 0)
	err := database.DB.Table("articles").Order("created_at DESC").Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

// 文章评论的展示
// 比较难，多理解，待优化（可能）
func GetArticleCommentsById(i, p int) (int64, []models.CommentVo, error) {
	var totle int64
	var commentslist []models.CommentVo
	err := database.DB.
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
		Find(&commentslist).Error
	if err != nil {
		return totle, commentslist, err
	}
	return totle, commentslist, nil
}

// 删除
func DeleteArticleById(id int) error {
	var art models.Article
	art.Id = id
	err := database.DB.Delete(&art).Error
	if err != nil {
		return err
	}
	return nil
}

//新增博客

func CreateArticle(article *models.CreatedOrUpdateArticleDto) error {

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

	err := database.DB.Table("articles").Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateArticleById(article *models.CreatedOrUpdateArticleDto) error {
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
	database.DB.Where("article_id = ?", article.Id).Delete(models.ArticleTag{})
	database.DB.Where("article_id = ?", article.Id).Delete(models.ArticleCategory{})
	for _, t := range tags {
		database.DB.Create(models.ArticleTag{ArticleId: article.Id, TagId: t.Id})
	}
	for _, c := range categories {
		database.DB.Create(models.ArticleCategory{ArticleId: article.Id, CategoryId: c.Id})
	}

	err := database.DB.Table("articles").Updates(
		models.Article{
			Id: article.Id, Title: article.Title, Description: article.Description,
			Content: article.Content, Status: article.Status, CreatedAt: article.CreatedAt, UpdatedAt: article.UpdatedAt,
		}).Where("id=?", article.Id).Error
	if err != nil {
		return err
	}
	return nil
}

func Search(pageNum, pageSize int, queryString string) ([]models.ArticleCategoryAndTag, int64, error) {

	var count int64
	list := make([]models.ArticleCategoryAndTag, 0)
	err := database.DB.Table("articles").
		Where(fmt.Sprintf("content like '%%%s'", queryString+"%")).
		Or(fmt.Sprintf("title like '%%%s'", queryString+"%")).
		Or(fmt.Sprintf("description like '%%%s'", queryString+"%")).
		Count(&count).
		Limit(pageSize).
		Offset(10 * (pageNum - 1)).
		Find(&list).Error
	if err != nil {
		return list, count, err
	}
	return list, count, nil
}
