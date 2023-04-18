package routes

import (
	"gin-blog/middleware"
	"gin-blog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouters() {
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(middleware.GinRecovery(true))
	r.Use(middleware.Cors())
	front := r.Group("api")
	{
		front.POST("login", useraboutApi.FrontLogin)
		front.GET("currentUser", useraboutApi.GetUserInfo)
		front.GET("auth/userInfo", middleware.AuthMiddleware(), useraboutApi.UserInfo)

		front.GET("overview", overviewApi.Overview)

		front.GET("articles", articleApi.ArticlePageNum)
		front.GET("articles/all", articleApi.GetArticlesAll)
		front.GET("articles/:id", articleApi.GetArticleDetailsById)
		front.GET("articles/:id/comments", articleApi.GetArticleComments)
		front.GET("articles/search", articleApi.SearchArticle)

		front.POST("comments", commentApi.Save)
		front.DELETE("comments/:id", commentApi.DelectComment)

		Oauth := front.Group("/oauth")
		{
			Oauth.GET("login/github", oauthApi.Login)
			Oauth.GET("callback/github", oauthApi.GithubLogincallback)
		}
	}
	v1 := r.Group("api")
	{
		v1.POST("auth/login", v1userApi.AuthLogin)

		v1.GET("comments/count", v1commentApi.CountCommentNotStatus)
		v1.GET("comments/", v1commentApi.GetCommentsByPage)
		v1.POST("comments/:id/audit/:status", v1commentApi.ChangeStatus)

		v1.GET("categories", v1categoryApi.GetCategoryALL)
		v1.POST("categories", v1categoryApi.AddCategory)
		v1.DELETE("categories/:id", v1categoryApi.DeleteCategory)

		v1.GET("tags", v1tagApi.GetTagALL)
		v1.POST("tags", v1tagApi.AddTag)
		v1.DELETE("tags/:id", v1tagApi.DeleteTag)

		v1.GET("users", v1userApi.GetUserList)
		v1.GET("users/:id", v1userApi.GetUserById)

		v1.DELETE("articles/:id", v1articleApi.DeleteArticle)
		v1.POST("articles", v1articleApi.AddArtilce)
		v1.PUT("articles/:id", v1articleApi.UpdateArtilce)

	}

	_ = r.Run(utils.HttpPort)
}
