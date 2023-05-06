package model

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Content string `gorm:"type:longtext; index:q_content_key,class:FULLTEXT,option:WITH PARSER ngram, not null"`

	AnswererID       uint // 外键索引
	AnswererNickName string

	QuestionID    uint // 外键索引
	QuestionTitle string

	Likes uint
}
