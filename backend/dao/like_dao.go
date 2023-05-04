package dao

import (
	"GoAsk/model"
	"errors"
	"gorm.io/gorm"
)

func (c *DbClient) GetAnswerLikes(aid uint) ([]uint, error) {
	var uids []uint
	err := c.Model(&model.Like{}).Where("answer_id = ?", aid).Pluck("user_id", &uids).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果 err 是没有找到记录，则不视为 error
			return []uint{}, nil
		} else {
			return []uint{}, err
		}
	}
	return uids, nil
}

// DeleteLike 删除更新
func (c *DbClient) DeleteLike(aid uint, uid uint) error {
	err := c.Where("answer_id = ? And user_id = ?", aid, uid).Delete(&model.Like{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果 err 是没有找到记录，则不视为 error
			return nil
		} else {
			return err
		}
	}
	// 因为取消点赞不会经常发生，所以每发生一次都对likes进行一次更新
	err = c.Exec("UPDATE answers SET likes = likes - 1 WHERE id = ?", aid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return err
}

// AddAnswerLikes 为某个 aid 批量添加点赞的用户
func (c *DbClient) AddAnswerLikes(aid uint, uids []uint) error {
	likes := make([]model.Like, len(uids))
	for i, _ := range likes {
		likes[i] = model.Like{AnswerID: aid, UserID: uids[i]}
	}
	err := c.Create(&likes).Error
	if err == nil { // 如果没出错，则执行
		err = c.Exec("UPDATE answers SET likes = likes + ?  WHERE id = ?;", len(uids), aid).Error
	}
	return err
}
