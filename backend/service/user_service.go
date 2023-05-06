package service

import (
	"GoAsk/dao"
	"GoAsk/serializer"
	"GoAsk/utils"
	e "GoAsk/utils/error"
	"context"
)

type UserService struct {
	UserID   uint
	UserName string `form:"user_name" json:"user_name" binding:"required,min=6,max=15"` // 限制，用户名最少6位，最多15位
	Password string `form:"password" json:"password" binding:"required,min=6,max=15"`   // 限制，密码最少6位，最多15位
	NickName string `form:"nick_name" json:"nick_name" binding:"max=10"`
}

type UserProfileService struct {
	UserID     uint
	Password   string `form:"password" json:"password"`
	NickName   string `form:"nick_name" json:"nick_name" binding:"max=10"` // 昵称，可以重复
	Profession string `form:"profession" json:"profession"`
	AboutMe    string `form:"about_me" json:"about_me"`
}

type GetUserService struct {
	UserID uint
}

func (service *UserService) Register(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	_, found, err := cli.GetUserByUserName(service.UserName)
	if err != nil {
		utils.Logger.Infoln(err)
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == true { // 该用户已经存在
		return serializer.BuildResponse(e.ErrorUserExist, nil)
	}
	passwordDigest, err := utils.EncryptPassword(service.Password)
	if err != nil {
		utils.Logger.Infoln(err)
		return serializer.BuildResponse(e.ErrorPasswordEncrypt, nil)
	}
	if service.NickName == "" {
		service.NickName = "匿名用户"
	}
	if user, err := cli.CreateUser(service.UserName, service.NickName, passwordDigest); err != nil {
		utils.Logger.Infoln(err)
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	} else { // 创建成功
		return serializer.BuildResponse(e.SUCCESS, serializer.BuildBriefUser(user))
	}
}

// Login 不使用set-cookies
func (service *UserService) Login(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	user, found, err := cli.GetUserByUserName(service.UserName)
	if found == false {
		if err == nil { // 用户不存在
			return serializer.BuildResponse(e.ErrorUserNotExist, nil)
		} else { // 数据库错误
			return serializer.BuildResponse(e.ErrorDataBase, nil)
		}
	} else {
		pass := utils.CheckPassword(service.Password, user.PasswordDigest)
		if pass == false { // 密码错误
			return serializer.BuildResponse(e.ErrorWrongPassword, nil)
		}
		tokenString, err := utils.GenerateToken(user.ID, user.UserName)
		if err != nil { // token签发错误
			utils.Logger.Infoln(err)
			return serializer.BuildResponse(e.ErrorTokenGenerate, nil)
		}
		// 捎带token给前端
		return serializer.BuildResponse(e.SUCCESS, serializer.BuildBriefUserWithToken(user, tokenString))
	}
}

//func (service *UserService) Login(ctx context.Context) (res serializer.Response, tokenString string) {
//	cli := dao.NewDbClient(ctx)
//	user, found, err := cli.GetUserByUserName(service.UserName)
//	if found == false {
//		if err == nil { // 用户不存在
//			return serializer.BuildResponse(e.ErrorUserNotExist, nil), ""
//		} else { // 数据库错误
//			return serializer.BuildResponse(e.ErrorDataBase, nil), ""
//		}
//	} else {
//		pass := utils.CheckPassword(service.Password, user.PasswordDigest)
//		if pass == false { // 密码错误
//			return serializer.BuildResponse(e.ErrorWrongPassword, nil), ""
//		}
//		tokenString, err := utils.GenerateToken(user.ID, user.UserName)
//		if err != nil { // token签发错误
//			utils.Logger.Infoln(err)
//			return serializer.BuildResponse(e.ErrorTokenGenerate, nil), ""
//		}
//		// 捎带token给前端
//		return serializer.BuildResponse(e.SUCCESS, serializer.BuildBriefUser(user)), tokenString
//	}
//}

// Update 更新用户信息（支持更新密码）
func (service *UserProfileService) Update(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	_, found, err := cli.GetUserByUserID(service.UserID)
	if found == false {
		if err == nil { // 用户不存在
			return serializer.BuildResponse(e.ErrorUserNotExist, nil)
		} else { // 数据库错误
			return serializer.BuildResponse(e.ErrorDataBase, nil)
		}
	} else {
		passwordDigest := ""
		if service.Password != "" {
			if len(service.Password) < 6 || len(service.Password) > 15 { // 下放到这里进行长度检验
				return serializer.BuildResponse(e.ErrorInvalidPasswordLength, nil)
			}
			passwordDigest, err = utils.EncryptPassword(service.Password)
			if err != nil { // 密码加密错误
				return serializer.BuildResponse(e.ErrorPasswordEncrypt, nil)
			}
		}
		user, err := cli.UpdateUser(service.UserID, passwordDigest, service.NickName,
			service.Profession, service.AboutMe)
		if err != nil { // 数据库错误
			return serializer.BuildResponse(e.ErrorDataBase, nil)
		}
		return serializer.BuildResponse(e.SUCCESS, serializer.BuildUser(user))
	}
}

func (service *GetUserService) Get(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	user, found, err := cli.GetUserByUserID(service.UserID)
	if found == false {
		if err == nil { // 用户不存在
			return serializer.BuildResponse(e.ErrorUserNotExist, nil)
		} else { // 数据库错误
			return serializer.BuildResponse(e.ErrorDataBase, nil)
		}
	} else {
		return serializer.BuildResponse(e.SUCCESS, serializer.BuildUser(user))
	}
}
