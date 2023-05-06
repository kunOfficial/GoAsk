package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	AnswerID uint //暂时不进行外键关联了
	UserID   uint
}
