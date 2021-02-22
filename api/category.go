package api

import (
	"github.com/kataras/iris/v12"
	"lookarm/model"
	"lookarm/utils/message"
)

//type QueryPagination struct {
//	pageSize int `url:"pagesize"`
//	pageNum int `url:"pagenum"`
//}

// 获取分类列表
func GetCategoryList(c iris.Context) {
	
	data, total, code := model.GetCategories()
	c.JSON(
		iris.Map{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": message.GetErrMsg(code),
		},
	)
}

// 获取单个分类
func GetCategory(c iris.Context) {
	id, _ := c.Params().GetInt("id")

	data, code := model.GetCategory(id)

	c.JSON(
		iris.Map{
			"status":  code,
			"data":    data,
			"message": message.GetErrMsg(code),
		},
	)
}

// 新增分类
func AddCategory(c iris.Context) {
	var data model.Category
	_ = c.ReadJSON(&data)
	code = model.CreateCategory(&data)
	c.JSON(
		iris.Map{
			"status":  code,
			"message": message.GetErrMsg(code),
		},
	)
}

// 编辑标签
func EditCategory(c iris.Context) {
	var data model.Category
	id, _ := c.Params().GetInt("id")
	_ = c.ReadJSON(&data)
	code = model.EditCategory(id, &data)

	c.JSON(
		iris.Map{
			"status":  code,
			"message": message.GetErrMsg(code),
		},
	)
}

// 删除标签
func DeleteCategory(c iris.Context) {
	id, _ := c.Params().GetInt("id")
	code = model.DeleteCategory(id)
	c.JSON(
		iris.Map{
			"status":  code,
			"message": message.GetErrMsg(code),
		},
	)
}
