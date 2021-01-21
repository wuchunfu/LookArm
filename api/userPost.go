package api

import (
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
	
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	
	if pageNum == 0 {
		pageNum = 1
	}
	
	data, total, code := model.GetPostInfoList(pageSize, pageNum)
	
	c.JSON(iris.Map{
		"data":    data,
		"total":   total,
		"status":  code,
		"message": message.GetErrMsg(code),
	})
}

// 查询单个表单


// 编辑表单

// 删除表单
