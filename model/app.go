package model

import (
	"gorm.io/gorm"
	"lookarm/utils/message"
)

type AppInfo struct {
	gorm.Model
	AppName      string `gorm:"type:varchar(200);not null" json:"app_name"`
	AppVersion   string `gorm:"type:varchar(200);not null" json:"app_version"`
	AppWebpage   string `gorm:"type:varchar(200);not null" json:"app_webpage"`
	AppDeveloper string `gorm:"type:varchar(200);not null" json:"app_developer"`
	AppDesc      string `gorm:"type:varchar(400);not null" json:"app_desc"`
	UserName     string `gorm:"type:varchar(200)" json:"user_name"`
	Email        string `gorm:"type:varchar(200)" json:"email"`
	CategoryID   int    `json:"category_id"`
	TagID        int    `json:"tag_id"`
	Category     Category
	Tag          Tag
}

// 查询app名是否重名
func CheckAppName(AppName string) int {
	var appInfo AppInfo
	db.Select("id").Where("app_name", AppName).First(&appInfo)
	if appInfo.ID > 0 {
		return message.ErrorAppExist
	}
	return message.SUCCESS
}

// 提交app信息
func AddAppInfo(data *AppInfo) {
	db.Create(&data)
}

// 查询app信息列表
func GetAppInfoList(pageSize int, pageNum int) ([]AppInfo, int64, int) {
	var appInfoList []AppInfo
	var total int64
	err = db.Joins("Category").Joins("Tag").Order("Updated_At DESC").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&appInfoList).Count(&total).Error
	if err != nil {
		return appInfoList, 0, message.ERROR
	}
	return appInfoList, total, message.SUCCESS
}

// 查询app信息
func SearchAppInfo(appName string, pageSize int, pageNum int) ([]AppInfo, int64, int) {
	var appInfoList []AppInfo
	var total int64
	err = db.Joins("Category").Joins("Tag").Order("Updated_At DESC").Where("app_name LIKE ?", "%"+appName+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&appInfoList).Count(&total).Error
	//db.Model(&appInfoList).Where("app_name LIKE ?", "%"+appName+"%")
	if err != nil {
		return appInfoList, 0, message.ERROR
	}
	return appInfoList, total, message.SUCCESS
}

// 查询分类下的App信息
func GetAppInfoCateList(cateID int, pageSize int, pageNum int) ([]AppInfo, int64, int) {
	var appInfoList []AppInfo
	var total int64

	err = db.Where("category_id = ?", cateID).Order("Updated_At DESC").Joins("Category").Joins("Tag").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&appInfoList).Count(&total).Error
	//db.Model(&appInfoList).Where("category_id = ?", cateID)
	if err != nil {
		return appInfoList, 0, message.ERROR
	}
	return appInfoList, total, message.SUCCESS
}

// 查询单个表单
func GetAppInfo(id int) (AppInfo, int) {
	var appInfo AppInfo
	err = db.Where("id = ?", id).First(&appInfo).Error
	if err != nil {
		return appInfo, message.ERROR
	}
	return appInfo, message.SUCCESS
}

// 编辑表单
func EditAppInfo(id int, data *AppInfo) int {
	err = db.Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}

// 删除表单
func DeleteAppInfo(id int) int {
	var appInfo AppInfo
	err = db.Where("id = ?", id).Delete(&appInfo).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCESS
}
