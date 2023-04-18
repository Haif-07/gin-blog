package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadZapLog(file)

}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8848")
	JwtKey = file.Section("server").Key("JwtKey").MustString("nihaoshijie")
}

func LoadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("myblog")
}
func LoadZapLog(file *ini.File) {
	Level = file.Section("zap-log").Key("Level").MustString("debug")
	Filename = file.Section("zap-log").Key("Filename").MustString("./log/test.log")
	MaxSize = file.Section("zap-log").Key("MaxSize").MustInt(1)
	MaxBackups = file.Section("zap-log").Key("MaxBackups").MustInt(5)
	MaxAge = file.Section("zap-log").Key("MaxAge").MustInt(30)
	Compress = file.Section("zap-log").Key("Compress").MustBool(false)

}
