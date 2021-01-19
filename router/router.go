package router

import (
	"github.com/kataras/iris/v12"
	"lookarm/api"
	"lookarm/middleware"
)

func InitRouter() {
	app := iris.Default()
	
	app.Use(middleware.Cors())
	
	v1 := app.Party("api/v1/")
	{
		// 管理员模块
		v1.Post("adduser", api.AddUser)
		v1.Delete("delete_user/{id}",api.DeleteUser)
		v1.Get("users",api.GetUsers)
	}
	
	_ = app.Listen(":8080")
	
}
