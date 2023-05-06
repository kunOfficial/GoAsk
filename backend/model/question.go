package model

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Title       string `gorm:"index:q_description_key,class:FULLTEXT,option:WITH PARSER ngram, not null"`
	Description string `gorm:"index:q_description_key,class:FULLTEXT,option:WITH PARSER ngram, not null"`
	ViewNum     uint
	// 用户的简要信息
	QuestionerID       uint // 外键索引
	QuestionerNickName string

	// 通过 has many 方式和 Answer 关联上
	Answers []Answer
}
