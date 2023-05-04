package dao

import (
	"GoAsk/model"
	"errors"
	"gorm.io/gorm"
)

// CreateQuestion 创建问题
func (c *DbClient) CreateQuestion(questionerID uint, title, description, questionerNickName string) (model.Question, error) {
	question := model.Question{
		Title:              title,
		Description:        description,
		ViewNum:            0,
		QuestionerID:       questionerID,
		QuestionerNickName: questionerNickName,
	}
	if err := c.Create(&question).Error; err != nil {
		return model.Question{}, err
	}
	return question, nil
}

// UpdateQuestion 更新问题
func (c *DbClient) UpdateQuestion(id uint, title, description string) (model.Question, error) {
	question := model.Question{
		Model:       gorm.Model{ID: id},
		Title:       title,
		Description: description,
	}
	err := c.Transaction(func(tx *gorm.DB) error { // 开启一个事务
		if err := c.Where("id=?", id).Updates(&question).Error; err != nil {
			return err
		}
		if title != "" { // 如果 title 改变了, 那该问题所有回答的 question_title 都要改变
			err := c.Model(&model.Answer{}).Where("question_id = ?", id).Update("question_title", title).Error
			if err != nil {
				return err
			}
		}
		if err := c.First(&question).Error; err != nil {
			return err
		} else {
			return nil
		}
	})
	if err == nil {
		return question, nil
	} else {
		return model.Question{}, err
	}
}

// AddQuestionView 更新浏览量
func (c *DbClient) AddQuestionView(id uint, viewNum uint) (found bool, err error) {
	// 更新放在一条 sql 语句中, 要不然就要使用事务来保证并发安全
	if err = c.Exec("UPDATE questions SET view_num = view_num + ? WHERE id = ?;", viewNum, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

// GetQuestionByID 获取某个问题(及其回答)
func (c *DbClient) GetQuestionByID(id uint, brief bool) (model.Question, bool, error) {
	var question model.Question
	question.ID = id
	var err error
	if brief {
		err = c.First(&question).Error // 不查回答的表
	} else {
		err = c.Preload("Answers").First(&question).Error // 预加载出所有回答
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
			return question, false, nil
		}
		return model.Question{}, false, err
	}
	return question, true, nil
}

// GetQuestions 获取所有问题
func (c *DbClient) GetQuestions(pageNum, pageSize uint) ([]model.Question, error) {
	var questions []model.Question
	err := c.Order("view_num DESC").
		Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize)).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// GetQuestionsByUserID 获取某个用户的所有问题
func (c *DbClient) GetQuestionsByUserID(pageNum, pageSize, userID uint) ([]model.Question, error) {
	var questions []model.Question
	err := c.Where("questioner_id=?", userID).
		Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize)).Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// GetQuestionsWithCond 根据筛选条件是否存在于(title, description)获取问题
func (c *DbClient) GetQuestionsWithCond(pageNum, pageSize uint, filter string) ([]model.Question, error) {
	var questions []model.Question
	err := c.Where("match(title, description) against(? in NATURAL LANGUAGE MODE)", filter).
		Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize)).
		Find(&questions).Error
	if err != nil {
		return nil, err
	}
	return questions, nil
}

// DeleteQuestion 删除某个问题
func (c *DbClient) DeleteQuestion(qid uint) (found bool, err error) {
	if err := c.Where("id=?", qid).Delete(&model.Question{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果 err 是没有找到记录，则不视为 error
			return false, nil
		} else {
			return false, err
		}
	}
	return true, err
}
