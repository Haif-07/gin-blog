package main

import (
	"fmt"
	"gin-blog/database"
	"gin-blog/log"
	"gin-blog/routes"
	"gin-blog/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(utils.Filename)
	gin.SetMode(utils.AppMode)
	err := log.InitLogger()
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	database.InitMysql()
	// r := gin.Default()
	routes.InitRouters()

}
