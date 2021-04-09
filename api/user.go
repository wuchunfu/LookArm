package api

import (
	"github.com/kataras/iris/v12"
	"lookarm/model"
	"lookarm/utils/message"
)

var code int

// AddUser 添加管理员
func AddUser(c iris.Context) {
	var data model.User
	_ = c.ReadJSON(&data)

	code = model.CreateUser(&data)

	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// GetUsers 查询管理员
func GetUsers(c iris.Context) {
	pageSize, _ := c.URLParamInt("pagesize")
	pageNum, _ := c.URLParamInt("pagenum")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	code, data, total := model.GetUsers(pageSize, pageNum)

	c.JSON(iris.Map{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// GetUserInfo 查询单个管理员
func GetUserInfo(c iris.Context) {
	id, _ := c.Params().GetInt("id")

	data, code := model.GetUserInfo(id)

	c.JSON(iris.Map{
		"data":    data,
		"status":  code,
		"message": message.GetErrMsg(code),
	})

}

// EditUser 编辑管理员
func EditUser(c iris.Context) {

	var data model.User
	data.ID, _ = c.Params().GetUint("id")

	_ = c.ReadJSON(&data)

	code = model.User.EditUserInfo(data)
	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// DeleteUser 删除管理员
func DeleteUser(c iris.Context) {
	var user model.User
	user.ID, _ = c.Params().GetUint("id")

	code = model.User.DeleteUser(user)

	c.JSON(iris.Map{"status": code, "message": message.GetErrMsg(code)})
}
