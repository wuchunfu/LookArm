package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"lookarm/model"
	"lookarm/utils/message"
)

// 提交表单
func PostAppInfo(c iris.Context) {
	var data model.PostInfo
	_ = c.ReadJSON(&data)

	code = model.PostAppInfo(&data)
	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 查询表单列表
func GetPostAppInfoList(c iris.Context) {
	pageSize, _ := c.URLParamInt("pagesize")
	pageNum, _ := c.URLParamInt("pagenum")
	appName := c.URLParam("app_name")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetPostInfoList(appName, pageSize, pageNum)

	c.JSON(iris.Map{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 查询分类下的表单
func GetPostInfoCateList(c iris.Context) {
	categoryId, _ := c.Params().GetInt("id")
	pageSize, _ := c.URLParamInt("pagesize")
	pageNum, _ := c.URLParamInt("pagenum")
	appName := c.URLParam("app_name")

	fmt.Println(categoryId)
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code := model.GetPostInfoCateList(categoryId, appName, pageSize, pageNum)

	c.JSON(iris.Map{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": message.GetErrMsg(code),
	})

}

// 查询单个表单
func GetPostInfo(c iris.Context) {

	id, _ := c.Params().GetInt("id")

	data, code := model.GetPostInfo(id)

	c.JSON(iris.Map{
		"data":    data,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 编辑表单
func EditPostInfo(c iris.Context) {
	var data model.PostInfo
	id, _ := c.Params().GetInt("id")
	_ = c.ReadJSON(&data)

	code = model.EditPostInfo(id, &data)

	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})

}

// 删除表单
func DeletePostInfo(c iris.Context) {
	id, _ := c.Params().GetInt("id")

	code = model.DeletePostInfo(id)
	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}
