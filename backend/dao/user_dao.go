package dao

import (
	"GoAsk/model"
	"GoAsk/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// GetUserByUserName 根据用户名搜索用户
func (c *DbClient) GetUserByUserName(userName string) (user model.User, found bool, err error) {
	if err := c.DB.Model(&model.User{}).Where("user_name=?", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 用户不存在
			return model.User{}, false, nil
		}
		// 其他问题
		return model.User{}, false, err
	}
	return user, true, nil
}

// CreateUser 保存用户
func (c *DbClient) CreateUser(userName, nickName, passwordDigest string) (user model.User, err error) {
	user.UserName = userName
	user.PasswordDigest = passwordDigest
	user.NickName = nickName
	if err := c.DB.Create(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}

// GetUserByUserID 获取某个User的信息
func (c *DbClient) GetUserByUserID(userID uint) (user model.User, found bool, err error) {
	if err := c.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 用户不存在
			return model.User{}, false, nil
		}
		// 其他问题
		return model.User{}, false, err
	}
	return user, true, nil
}

// UpdateUser 更新用户信息
func (c *DbClient) UpdateUser(userID uint, passwordDigest, nickName, profession, aboutMe string) (model.User, error) {
	user := model.User{
		Model:          gorm.Model{ID: userID},
		NickName:       nickName,
		Profession:     profession,
		AboutMe:        aboutMe,
		PasswordDigest: passwordDigest,
	}
	if err := c.Where("id=?", user.ID).Updates(user).Error; err != nil {
		utils.Logger.Infoln(err)
		return model.User{}, err
	}
	if err := c.First(&user).Error; err != nil {
		return model.User{}, err
	}
	return user, nil
}
