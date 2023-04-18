package database

import (
	"fmt"
	"gin-blog/utils"
	"os"
	"time"

	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// v2版本的GORM不用手动关闭了，https://github.com/go-gorm/gorm/issues/3834
func InitMysql() {
	//dsn := "root:123456@/myblog?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("连接数据库失败，请检查参数：", zap.Error(err))
		os.Exit(1)
	}
	DB = db
	sqlDB, err := db.DB()
	if err != nil {
		zap.L().Error("获取通用数据库对象出错：", zap.Error(err))
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}
