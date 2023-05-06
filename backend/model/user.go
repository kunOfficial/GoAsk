package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PasswordDigest string `gorm:"not null"`
	UserName       string `gorm:"not null"` // 用户名，用于登录，不可重复
	NickName       string // 昵称，可以重复
	Profession     string
	AboutMe        string
	//// 通过 has many 关系关联上 Question 和 Answer
	//Questions []Question `gorm:"foreignKey:QuestionerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	//Answers   []Answer   `gorm:"foreignKey:AnswererID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
