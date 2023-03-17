package routes

import (
	"gin-blog/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
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

	return r
}

// func AboutUser(e *gin.Engine) {
// 	User := e.Group("/api")

// 	{
// 		User.POST("login", useraboutApi.FrontLogin)
// 		User.GET("currentUser", useraboutApi.GetUserInfo)
// 		User.GET("auth/userInfo", middleware.AuthMiddleware(), useraboutApi.UserInfo)
// 	}
// }

// func Auth(e *gin.Engine) {
// 	Oauth := e.Group("/api/oauth")

// 	{
// 		Oauth.GET("login/github", oauthApi.Login)
// 		Oauth.GET("callback/github", oauthApi.GithubLogincallback)
// 	}
// }

// func Overview(e *gin.Engine) {
// 	Overview := e.Group("/api")

// 	{
// 		Overview.GET("overview", overviewApi.Overview)
// 	}
// }

// func Article(e *gin.Engine) {
// 	FrontArticle := e.Group("/api")

// 	{
// 		FrontArticle.GET("articles", articleApi.ArticlePageNum)
// 		FrontArticle.GET("articles/all", articleApi.GetArticlesAll)
// 		FrontArticle.GET("articles/:id", articleApi.GetArticleDetailsById)
// 		FrontArticle.GET("articles/:id/comments", articleApi.GetArticleComments)

// 	}
// }

// func Comment(e *gin.Engine) {
// 	Comment := e.Group("/api")

// 	{
// 		Comment.POST("comments", commentApi.Save)
// 		Comment.DELETE("comments/:id", commentApi.DelectById)
// 	}
// }
