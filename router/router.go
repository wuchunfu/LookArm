package router

import (
	"github.com/kataras/iris/v12"
	"lookarm/api"
	"lookarm/middleware"
)

func InitRouter() {
	app := iris.Default()
	
	app.Use(middleware.Cors())
	
	app.Post("/adduser",api.AddUser)
	
	_ = app.Listen(":8080")
	
}
