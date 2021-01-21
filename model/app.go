package model

import (
	"gorm.io/gorm"
	"lookarm/utils/message"
)

type AppInfo struct {
	gorm.Model
	AppName      string   `gorm:"type:varchar(200);not null" json:"app_name"`
	AppVersion   string   `gorm:"type:varchar(200);not null" json:"app_version"`
	AppWebpage   string   `gorm:"type:varchar(200);not null" json:"app_webpage"`
	AppDeveloper string   `gorm:"type:varchar(200);not null" json:"app_developer"`
	AppDesc      string   `gorm:"type:varchar(400);not null" json:"app_desc"`
	UserName     string   `gorm:"type:varchar(200)" json:"user_name"`
	Email        string   `gorm:"type:varchar(200)" json:"email"`
	CategoryID   int      `json:"category_id"`
	TagID        int      `json:"tag_id"`
	Category     Category `gorm:"foreignKey:CategoryID"`
	Tag          Tag      `gorm:"foreignKey:TagID"`
}

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
func GetPostInfoList(pageSize int, pageNum int) ([]PostInfo, int64, int) {
	var postInfoList []PostInfo
	var total int64

	err = db.Preload("Category").Preload("Tag").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&postInfoList).Error
	db.Model(&postInfoList).Count(&total)
	if err != nil {
		return postInfoList, 0, message.ERROR
	}
	return postInfoList, total, message.SUCCSES
}

// 查询单个表单
func GetPostInfo(id int) (PostInfo, int) {
	var postInfo PostInfo
	err = db.Preload("category").Preload("Tag").Where("id = ?", id).First(&postInfo).Error
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

// 以下是app信息的接口模块
// 思路：

// 提交app信息
func AddAppInfo(data *AppInfo) int {
	err = db.Create(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}

// 查询app信息列表
func GetAppInfoList(pageSize int, pageNum int) ([]AppInfo, int64, int) {
	var appInfoList []AppInfo
	var total int64

	err = db.Preload("Category").Preload("Tag").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&appInfoList).Error
	db.Model(&appInfoList).Count(&total)
	if err != nil {
		return appInfoList, 0, message.ERROR
	}
	return appInfoList, total, message.SUCCSES
}

// 查询单个表单
func GetAppInfo(id int) (AppInfo, int) {
	var appInfo AppInfo
	err = db.Preload("Category").Preload("Tag").Where("id = ?", id).First(&appInfo).Error
	if err != nil {
		return appInfo, message.ERROR
	}
	return appInfo, message.SUCCSES
}

// 编辑表单
func EditAppInfo(id int, data *AppInfo) int {
	err = db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}

// 删除表单
func DeleteAppInfo(id int) int {
	var appInfo AppInfo
	err = db.Where("id = ?", id).Delete(&appInfo).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}
