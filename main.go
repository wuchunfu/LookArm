package main

import (
	"lookarm/model"
	"lookarm/router"
)

func main()  {
	// 引入数据库
	model.InitDatabase()
	
	// 初始化路由
	router.InitRouter()
}
