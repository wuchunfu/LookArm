package router

import (
	"github.com/kataras/iris/v12"
)

func InitRouter()  {
	app := iris.Default()
	
	app.Get("/hello", func(c iris.Context) {
		c.JSON(iris.Map{"hello":"weject"})
	})
	
	app.Listen(":8080")
	
}
