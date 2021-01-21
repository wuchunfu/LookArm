package api

import (
	"github.com/kataras/iris/v12"
	"lookarm/middleware"
	"lookarm/model"
	"lookarm/utils/message"
)

// 登录

func Login(c iris.Context) {
	var loginData model.User
	var token string
	_ = c.ReadJSON(&loginData)
	
	loginData, code = model.CheckLogin(loginData.UserName, loginData.Password)
	
	if code == message.SUCCSES {
		token, code = middleware.SetToken(loginData.UserName)
	}
	c.JSON(iris.Map{
		"status":  code,
		"data":    loginData.UserName,
		"id":      loginData.ID,
		"message": message.GetErrMsg(code),
		"token":   token,
	})
	
}
