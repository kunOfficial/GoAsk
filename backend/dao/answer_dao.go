package dao

import (
	"GoAsk/model"
	"errors"
	"gorm.io/gorm"
)

func (c *DbClient) CreateAnswer(answererID, questionID uint, content, questionTitle, answererNickName string) (model.Answer, error) {
	answer := model.Answer{
		Content:          content,
		AnswererID:       answererID,
		QuestionID:       questionID,
		QuestionTitle:    questionTitle,
		AnswererNickName: answererNickName,
	}
	if err := c.Create(&answer).Error; err != nil {
		return model.Answer{}, err
	}
	return answer, nil
}

func (c *DbClient) GetAnswerByID(id uint) (model.Answer, bool, error) {
	var answer model.Answer
	answer.ID = id
	err := c.First(&answer).Error
	if err == nil {
		return answer, true, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) { // 无记录
		return model.Answer{}, false, nil
	} else {
		return model.Answer{}, false, err
	}

}

func (c *DbClient) GetAnswers(pageNum, pageSize uint) ([]model.Answer, error) {
	var answers []model.Answer
	err := c.Order("likes DESC").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&answers).Error
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (c *DbClient) GetAnswersByUserID(pageNum, pageSize, UserID uint) ([]model.Answer, error) {
	var answers []model.Answer
	err := c.Limit(int(pageSize)).Offset(int((pageNum-1)*pageSize)).Where("answerer_id = ?", UserID).Find(&answers).Error
	if err != nil {
		return nil, err
	}
	return answers, nil
}

// DeleteAnswer 删除某个问题
func (c *DbClient) DeleteAnswer(aid uint) (found bool, err error) {
	if err := c.Where("id=?", aid).Delete(&model.Answer{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果 err 是没有找到记录，则不视为 error
			return false, nil
		} else {
			return false, err
		}
	}
	return true, err
}

// UpdateAnswer 更新回答
func (c *DbClient) UpdateAnswer(id uint, content string) (model.Answer, error) {
	answer := model.Answer{
		Model:   gorm.Model{ID: id},
		Content: content,
	}
	if err := c.Where("id=?", id).Updates(&answer).Error; err != nil {
		return model.Answer{}, err
	}
	if err := c.First(&answer).Error; err != nil {
		return model.Answer{}, err
	} else {
		return answer, nil
	}
}
