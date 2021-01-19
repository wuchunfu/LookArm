package model

import "gorm.io/gorm"

type AppInfo struct {
	gorm.Model
	AppName     string   `gorm:"type:varchar(200);not null" json:"app_name"`
	AppVersion  string   `gorm:"type:varchar(200);not null" json:"app_version"`
	AppWebpage  string   `gorm:"type:varchar(200);not null" json:"app_webpage"`
	AppDesc     string   `gorm:"type:varchar(400);not null" json:"app_desc"`
	UserName    string   `gorm:"type:varchar(200)" json:"user_name"`
	Email       string   `gorm:"type:varchar(200)" json:"email"`
	AppCategory int      `gorm:"not null" json:"app_category"`
	AppTag      int      `gorm:"not null" json:"app_tag"`
	Category    Category `gorm:"foreignKey:AppCategory"`
	Tag         Tag      `gorm:"foreignKey:AppTag"`
}

type PostInfo struct {
	gorm.Model
	AppName     string   `gorm:"type:varchar(200);not null" json:"app_name"`
	AppVersion  string   `gorm:"type:varchar(200);not null" json:"app_version"`
	AppWebpage  string   `gorm:"type:varchar(200);not null" json:"app_webpage"`
	AppDesc     string   `gorm:"type:varchar(400);not null" json:"app_desc"`
	UserName    string   `gorm:"type:varchar(200)" json:"user_name"`
	Email       string   `gorm:"type:varchar(200)" json:"email"`
	AppCategory int      `gorm:"not null" json:"app_category"`
	AppTag      int      `gorm:"not null" json:"app_tag"`
	Category    Category `gorm:"foreignKey:AppCategory"`
	Tag         Tag      `gorm:"foreignKey:AppTag"`
}
