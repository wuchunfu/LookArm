package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"lookarm/model"
	"lookarm/utils/message"
)

// 提交App信息
func AddAppInfo(c iris.Context) {
	var data model.AppInfo
	_ = c.ReadJSON(&data)
	
	code = model.CheckAppName(data.AppName)
	fmt.Println(data.AppName)
	if code == message.SUCCESS {
		model.AddAppInfo(&data)
	}
	
	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 查询表单列表
func GetAppInfoList(c iris.Context) {
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
	
	data, total, code := model.GetAppInfoList(pageSize, pageNum)
	
	c.JSON(iris.Map{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 搜索App信息
func SearchAppInfo(c iris.Context) {
	var data []model.AppInfo
	var total int64
	pageSize, _ := c.URLParamInt("pagesize")
	pageNum, _ := c.URLParamInt("pagenum")
	appName := c.URLParam("app_name")
	
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	
	if pageNum <= 0 {
		pageNum = 1
	}
	
	if appName == "null" {
		data, total, code = model.GetAppInfoList(pageSize, pageNum)
		c.JSON(iris.Map{
			"data":    data,
			"total":   total,
			"status":  code,
			"message": message.GetErrMsg(code),
		})
		return
	}
	
	data, total, code = model.SearchAppInfo(appName, pageSize, pageNum)
	c.JSON(iris.Map{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 查询分类下的APP信息
func GetAppInfoCateList(c iris.Context) {
	categoryId, _ := c.Params().GetInt("id")
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
	
	data, total, code := model.GetAppInfoCateList(categoryId, pageSize, pageNum)
	
	c.JSON(iris.Map{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
	
}

// 查询单个App信息
func GetAppInfo(c iris.Context) {
	
	id, _ := c.Params().GetInt("id")
	
	data, code := model.GetAppInfo(id)
	
	c.JSON(iris.Map{
		"data":    data,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 编辑表单
func EditAppInfo(c iris.Context) {
	var data model.AppInfo
	
	id, _ := c.Params().GetInt("id")
	_ = c.ReadJSON(&data)
	
	code = model.EditAppInfo(id, &data)
	
	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})
	
}

// 删除表单
func DeleteAppInfo(c iris.Context) {
	id, _ := c.Params().GetInt("id")
	
	code = model.DeleteAppInfo(id)
	c.JSON(iris.Map{
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}
