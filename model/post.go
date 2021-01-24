package model

import (
	"gorm.io/gorm"
	"lookarm/utils/message"
)

type PostInfo struct {
	gorm.Model
	AppName      string   `gorm:"type:varchar(200);not null" json:"app_name"`
	AppVersion   string   `gorm:"type:varchar(200);not null" json:"app_version"`
	AppWebpage   string   `gorm:"type:varchar(200);not null" json:"app_webpage"`
	AppDesc      string   `gorm:"type:varchar(400);not null" json:"app_desc"`
	AppDeveloper string   `gorm:"type:varchar(200);not null" json:"app_developer"`
	UserName     string   `gorm:"type:varchar(200)" json:"user_name"`
	Email        string   `gorm:"type:varchar(200)" json:"email"`
	CategoryID   int      `json:"category_id"`
	TagID        int      `json:"tag_id"`
	Category     Category `gorm:"foreignKey:CategoryID"`
	Tag          Tag      `gorm:"foreignKey:TagID"`
}

// 提交表单
func PostAppInfo(data *PostInfo) int {
	err = db.Create(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}

// 查询表单列表
func GetPostInfoList(appName string, pageSize int, pageNum int) ([]PostInfo, int64, int) {
	var postInfoList []PostInfo
	var total int64
	if appName == "" {
		err = db.Preload("Category").Preload("Tag").Order("Updated_At DESC").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&postInfoList).Error
		db.Model(&postInfoList).Count(&total)
	}
	err = db.Preload("Category").Order("Updated_At DESC").Preload("Tag").Where("app_name LIKE ?", "%"+appName+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&postInfoList).Error
	db.Model(&postInfoList).Where("app_name like ?", "%"+appName+"%").Count(&total)
	if err != nil {
		return postInfoList, 0, message.ERROR
	}
	return postInfoList, total, message.SUCCSES
}

// 查询分类下的表单
func GetPostInfoCateList(cateID int, appName string, pageSize int, pageNum int) ([]PostInfo, int64, int) {
	var postInfoList []PostInfo
	var total int64
	if appName == "" {
		err = db.Where("category_id = ?", cateID).Order("Updated_At DESC").Preload("Category").Preload("Tag").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&postInfoList).Error
		db.Model(&postInfoList).Where("category_id = ?", cateID).Count(&total)
	}
	err = db.Where("category_id = ?", cateID).Where("app_name LIKE ?", "%"+appName+"%").Order("Updated_At DESC").Preload("Category").Preload("Tag").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&postInfoList).Error
	db.Model(&postInfoList).Where("app_name like ?", "%"+appName+"%").Where("category_id = ?", cateID).Count(&total)
	if err != nil {
		return postInfoList, 0, message.ERROR
	}
	return postInfoList, total, message.SUCCSES
}

// 查询单个表单
func GetPostInfo(id int) (PostInfo, int) {
	var postInfo PostInfo
	err = db.Where("id = ?", id).First(&postInfo).Error
	if err != nil {
		return postInfo, message.ERROR
	}
	return postInfo, message.SUCCSES
}

// 编辑表单
func EditPostInfo(id int, data *PostInfo) int {
	err = db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}

// 删除表单
func DeletePostInfo(id int) int {
	var postInfo PostInfo
	err = db.Where("id = ?", id).Delete(&postInfo).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}