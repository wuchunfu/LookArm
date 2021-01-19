package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"lookarm/utils/message"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(200);not null" json:"user_name"`
	Password string `gorm:"type:varchar(200);not null" json:"password"`
	Email    string `gorm:"type:varchar(200);not null" json:"email"`
}

// 增加管理员
func (u User) CreateUser() int {
	err = db.Create(&u).Error
	if err != nil {
		return message.ERROR
	}
	return message.SUCCSES
}

// 查询用户列表
func (u User) GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err = db.Select("id,user_name,email").Limit(pageSize).Offset((pageNum -1) * pageSize).Find(&users).Error
	db.Model(&u).Count(&total)
	if err != nil {
		return users,0
	}
	return users,total
}

// 编辑用户
func (u User) EditUserInfo(id int)  {

}

// 生成密码
func SetPassword(password string) string {
	const cost = 10
	
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatal(err)
	}
	return string(HashPw)
}
