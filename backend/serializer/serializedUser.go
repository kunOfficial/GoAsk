package serializer

import (
	"GoAsk/model"
)

type SerializedBriefUser struct {
	ID        uint   `json:"user_id" form:"user_id" example:"1"`
	NickName  string `json:"nick_name" form:"nick_name" example:"FanOne"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

type SerializedUserWithToken struct {
	SerializedBriefUser
	Token string `json:"token" form:"token"`
}

type SerializedUser struct {
	SerializedBriefUser
	UserName   string `json:"user_name"`
	AboutMe    string `json:"about_me"`
	Profession string `json:"profession"`
}

// BuildBriefUser 序列化用户
func BuildBriefUser(user model.User) SerializedBriefUser {
	return SerializedBriefUser{
		ID:        user.ID,
		NickName:  user.NickName,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func BuildUser(user model.User) SerializedUser {
	return SerializedUser{
		SerializedBriefUser: SerializedBriefUser{
			ID:        user.ID,
			NickName:  user.NickName,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		},
		UserName:   user.UserName,
		AboutMe:    user.AboutMe,
		Profession: user.Profession,
	}
}

func BuildBriefUserWithToken(user model.User, token string) SerializedUserWithToken {
	return SerializedUserWithToken{
		SerializedBriefUser: BuildBriefUser(user),
		Token:               token,
	}
}
