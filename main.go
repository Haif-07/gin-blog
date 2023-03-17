package main

import (
	"fmt"
	databas "gin-blog/database"
	"gin-blog/routes"
)

func main() {

	err := databas.InitMysql()
	if err != nil {
		fmt.Println("数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
	}
	// r := gin.Default()
	// routes.AboutUser(r)
	// routes.Auth(r)
	// routes.Overview(r)
	// routes.Article(r)
	// routes.Comment(r)
	e := routes.InitRouters()
	e.Run(":8848")
	defer databas.Close()

}
