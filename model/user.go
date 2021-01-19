package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(200);not null" json:"user_name"`
	Password string `gorm:"type:varchar(200);not null" json:"password"`
	Email    string `gorm:"type:varchar(200);not null" json:"email"`
}
