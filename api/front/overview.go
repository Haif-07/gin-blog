package front

import (
	"gin-blog/dao"
	"gin-blog/models"

	"github.com/gin-gonic/gin"
)

type Overview struct{}
type ResourceCount struct {
	Article  int `json:"article"`
	Comment  int `json:"comment"`
	Category int `json:"category"`
	Tag      int `json:"tag"`
}

type OverviewVo struct {
	ResourceCount ResourceCount        `json:"resourceCount"`
	Tags          []models.Tagdto      `json:"tags"`
	Categories    []models.Categorydto `json:"categories"`
	Pages         []models.Page        `json:"pages"`
	Configs       []models.Config      `json:"configs"`
}

func (*Overview) Overview(c *gin.Context) {
	var o OverviewVo
	o.ResourceCount.Article = dao.GetArticleCount()
	o.ResourceCount.Category = dao.GetCategoryCount()
	o.ResourceCount.Comment = dao.GetCommentCount()
	o.ResourceCount.Tag = dao.GetTagCount()
	o.Tags = dao.GetTagsWithArticleCount()
	o.Categories = dao.GetCategoryWithArticleCount()
	o.Pages = dao.GetPage()
	o.Configs = dao.Getconfigs()
	// o.ResourceCount.Article = 10
	// o.ResourceCount.Category = 3
	// o.ResourceCount.Comment = 4
	// o.ResourceCount.Tag = 5
	// tag1 := Tag{
	// 	Id:           7,
	// 	Name:         "软件过程管理",
	// 	ArticleCount: 4,
	// 	CreatedAt:    "2020-12-18 14:55",
	// 	UpdatedAt:    "2020-12-18 14:55",
	// }
	// tag2 := Tag{
	// 	Id:           8,
	// 	Name:         "读书笔记",
	// 	ArticleCount: 4,
	// 	CreatedAt:    "2020-12-18 14:55",
	// 	UpdatedAt:    "2020-12-18 14:55",
	// }
	// categories1 := Categories{
	// 	Id:           6,
	// 	Name:         "读书笔记",
	// 	ArticleCount: 4,
	// 	CreatedAt:    "2020-12-18 14:55",
	// 	UpdatedAt:    "2020-12-18 14:55",
	// }
	// categories2 := Categories{
	// 	Id:           7,
	// 	Name:         "嘻嘻哈啊哈",
	// 	ArticleCount: 4,
	// 	CreatedAt:    "2020-12-18 14:55",
	// 	UpdatedAt:    "2020-12-18 14:55",
	// }
	// pages1 := Pages{
	// 	Id:        3,
	// 	Name:      "关于",
	// 	Icon:      "user",
	// 	Link:      "/about",
	// 	Status:    1,
	// 	CreatedAt: "2020-12-18 14:55",
	// 	UpdatedAt: "2020-12-18 14:55",
	// }
	// pages2 := Pages{
	// 	Id:        43,
	// 	Name:      "nihao",
	// 	Icon:      "user",
	// 	Link:      "/about",
	// 	Status:    2,
	// 	CreatedAt: "2020-12-18 14:55",
	// 	UpdatedAt: "2020-12-18 14:55",
	// }

	// config1 := Configs{
	// 	Id:    1,
	// 	Name:  "indexTitle",
	// 	Value: "跑起来就有风",
	// }
	// config2 := Configs{
	// 	Id:    2,
	// 	Name:  "bannerTitle",
	// 	Value: "跑起来就有风",
	// }
	// o.Tag = append(o.Tag, tag1)
	// o.Tag = append(o.Tag, tag2)
	// o.Categories = append(o.Categories, categories1)
	// o.Categories = append(o.Categories, categories2)
	// o.Pages = append(o.Pages, pages1)
	// o.Pages = append(o.Pages, pages2)

	// o.Configs = append(o.Configs, config1)
	// o.Configs = append(o.Configs, config2)

	c.JSON(200, gin.H{
		"code":    200000,
		"message": "success",
		"data":    o,
	})
}
