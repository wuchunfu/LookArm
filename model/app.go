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
	AppCategory  int      `gorm:"not null" json:"app_category"`
	AppTag       int      `gorm:"not null" json:"app_tag"`
	Category     Category `gorm:"foreignKey:AppCategory"`
	Tag          Tag      `gorm:"foreignKey:AppTag"`
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
	AppCategory  int      `gorm:"not null" json:"app_category"`
	AppTag       int      `gorm:"not null" json:"app_tag"`
	Category     Category `gorm:"foreignKey:AppCategory"`
	Tag          Tag      `gorm:"foreignKey:AppTag"`
}

// 提交表单
func PostAppInfo(data *PostInfo) int {
	err = db.Create(&data).Error
	if err != nil {
		return message.ERROR
	}
	return  message.SUCCSES
}

// 查询表单列表
func GetPostInfoList(pageSize int,pageNum int) ([]PostInfo,int64,int) {
	var postInfoList []PostInfo
	var total int64
	
	err = db.Preload("Category","Tag").Limit(pageSize).Offset((pageSize-1)*pageNum).Find(&postInfoList).Count(&total).Error
	if err != nil {
		return postInfoList,0,message.ERROR
	}
	return postInfoList,total,message.SUCCSES
}

// 查询单个表单
func GetPostInfo(id int) (PostInfo,int) {
	var postInfo PostInfo
	err = db.Preload("Category","Tag").Where("id = ?",id).First(&postInfo).Error
	if err != nil {
		return postInfo,message.ERROR
	}
	return postInfo,message.SUCCSES
}

// 编辑表单
func EditPostInfo(id int,data *PostInfo) int {
	err = db.Where("id = ?",id).Updates(&data).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}

// 删除表单
func DeletePostInfo(id int) int {
	var postInfo PostInfo
	err = db.Where("id = ?",id).Delete(&postInfo).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}