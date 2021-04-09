package model

import "lookarm/utils/message"

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(200)" json:"name"`
}

// GetCategories 获取分类列表
func GetCategories(pageSize int, pageNum int) ([]Category, int64, int) {
	var cates []Category
	var total int64
	err = db.Select("id","name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	db.Model(&cates).Count(&total)
	if err != nil {
		return cates, 0, message.ERROR
	}

	return cates, total, message.SUCCESS
}

// GetCategory 获取标签
func GetCategory(id int) (Category, int) {
	var cate Category
	err = db.Where("id = ?", id).Limit(1).Find(&cate).Error
	if err != nil {
		return cate, message.ERROR
	}
	return cate, message.SUCCESS
}

// CreateCategory 新增标签
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
func DeleteCategory(id int) int {
	var cate Category
	err = db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}
