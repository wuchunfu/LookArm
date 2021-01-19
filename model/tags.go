package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	TagName string `gorm:"type:varchar(200)" json:"tag_name"`
}
