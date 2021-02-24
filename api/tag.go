package api

import (
	"github.com/kataras/iris/v12"
	"lookarm/model"
	"lookarm/utils/message"
)

// 获取标签列表
func GetTagList(c iris.Context) {
	pageSize := c.URLParamIntDefault("pagesize",10)
	pageNum := c.URLParamIntDefault("pagenum",1)
	
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	
	if pageNum == 0 {
		pageNum = 1
	}
	
	data, total, code := model.GetTags(pageSize, pageNum)
	c.JSON(
		iris.Map{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": message.GetErrMsg(code),
		},
	)
}

// 获取单个标签
func GetTagInfo(c iris.Context) {
	id, _ := c.Params().GetInt("id")
	
	data, code := model.GetTag(id)
	
	c.JSON(
		iris.Map{
			"status":  code,
			"data":    data,
			"message": message.GetErrMsg(code),
		},
	)
}

// 新增标签
func AddTag(c iris.Context) {
	var data model.Tag
	_ = c.ReadJSON(&data)
	code = model.CreateTag(&data)
	c.JSON(
		iris.Map{
			"status":  code,
			"message": message.GetErrMsg(code),
		},
	)
}

// 编辑标签
func EditTag(c iris.Context) {
	var data model.Tag
	id, _ := c.Params().GetInt("id")
	_ = c.ReadJSON(&data)
	code = model.EditTag(id, &data)
	
	c.JSON(
		iris.Map{
			"status":  code,
			"message": message.GetErrMsg(code),
		},
	)
}

// 删除标签
func DeleteTag(c iris.Context) {
	id, _ := c.Params().GetInt("id")
	code = model.DeteleTag(id)
	c.JSON(
		iris.Map{
			"status":  code,
			"message": message.GetErrMsg(code),
		},
	)
}
