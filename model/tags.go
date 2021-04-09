package model

import (
	"lookarm/utils/message"
)

type Tag struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	TagName string `gorm:"type:varchar(200)" json:"tag_name"`
}

// GetTags 获取标签列表
func GetTags(pageSize int, pageNum int) ([]Tag, int64, int) {
	var tags []Tag
	var total int64
	err = db.Select("id","tag_name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&tags).Error
	db.Model(&tags).Count(&total)
	if err != nil {
		return nil, 0, message.ERROR
	}
	return tags, total, message.SUCCESS
}

// GetTag 获取标签
func GetTag(id int) (Tag, int) {
	var tag Tag
	err = db.Where("id = ?", id).Limit(1).Find(&tag).Error
	if err != nil {
		return tag, message.ERROR
	}
	return tag, message.SUCCESS

}

// CreateTag 新增标签
func CreateTag(data *Tag) int {
	err := db.Create(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}

// EditTag 修改标签
func EditTag(id int, data *Tag) int {
	err = db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}

// DeleteTag 删除标签
func DeleteTag(id int) int {
	var tag Tag
	err = db.Where("id = ?", id).Delete(&tag).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}
