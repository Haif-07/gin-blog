package routes

import (
	"gin-blog/api/front"
	v1 "gin-blog/api/v1"
)

var (

	//前端请求
	useraboutApi front.User
	oauthApi     front.Oauth
	overviewApi  front.Overview
	articleApi   front.Article
	commentApi   front.Comment
	//后端请求

	v1userApi     v1.User
	v1commentApi  v1.Comment
	v1categoryApi v1.Category
	v1tagApi      v1.Tag
	v1articleApi  v1.Article
)
