package api

import (
	"github.com/kataras/iris/v12"
	"lookarm/model"
	"lookarm/utils/message"
)

var code int

// 添加管理员
func AddUser(c iris.Context)  {
	var data model.User
	c.ReadJSON(&data)
	
	code = model.User.CreateUser(data)
	
	c.JSON(iris.Map{
		"status":code,
		"message":message.GetErrMsg(code),
	})
}
