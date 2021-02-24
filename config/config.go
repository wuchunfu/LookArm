package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	ServerPort string
	JwtKey     string
	//LogLevel string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	ServerPort = file.Section("Server").Key("serverPort").String()
	JwtKey = file.Section("Server").Key("jwtKey").String()
	//LogLevel = file.Section("Server").Key("logLevel").String()
}

func LoadData(file *ini.File) {
	DbHost = file.Section("Database").Key("db_host").String()
	DbPort = file.Section("Database").Key("db_port").String()
	DbUser = file.Section("Database").Key("db_user").String()
	DbPassword = file.Section("Database").Key("db_password").String()
	DbName = file.Section("Database").Key("db_name").String()
}
