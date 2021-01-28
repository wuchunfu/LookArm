package model

import "lookarm/utils/message"

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(200)" json:"name"`
}

// 获取分类列表
func GetCategories(pageSize int, pageNum int) ([]Category, int64, int) {
	var cates []Category
	var total int64
	err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&cates).Count(&total).Error
	if err != nil {
		return cates, 0, message.ERROR
	}

	return cates, total, message.SUCCESS
}

// 获取标签
func GetCategory(id int) (Category, int) {
	var cate Category
	err = db.Where("id = ?", id).First(&cate).Error
	if err != nil {
		return cate, message.ERROR
	}
	return cate, message.SUCCESS

}

// 新增标签
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}

// 修改标签
func EditCategory(id int, data *Category) int {
	err = db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}

// 删除标签
func DeteleCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}
