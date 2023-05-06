package v1

import (
	"GoAsk/serializer"
	"GoAsk/service"
	e "GoAsk/utils/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegisterService service.UserService
	if err := c.ShouldBind(&userRegisterService); err != nil { // 这里通过shouldBind进行绑定，可以将值传到userRegister变量中
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
	} else {
		res := userRegisterService.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// 采用 cookies 验证
//func UserLogin(c *gin.Context) {
//	var userLoginService service.UserService
//	if err := c.ShouldBind(&userLoginService); err != nil { // 这里通过shouldBind进行绑定，可以将值传到userRegister变量中
//		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
//	} else {
//		res, tokenString := userLoginService.Login(c.Request.Context())
//		if tokenString != "" {
//			// secure 是 false, 即不使用https
//			c.SetCookie("token", tokenString, 24*60*60, "/",
//				"localhost", false, true)
//		}
//		c.JSON(http.StatusOK, res)
//	}
//}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err != nil { // 这里通过shouldBind进行绑定，可以将值传到userRegister变量中
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
	} else {
		res := userLoginService.Login(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// UpdateUserProfile 更新用户信息
func UpdateUserProfile(c *gin.Context) {
	var updateUserProfileService service.UserProfileService
	userID := c.MustGet("user_id").(uint)
	if id, _ := strconv.Atoi(c.Param("id")); id != int(userID) { // 检验一下是否UserID是否对得上
		c.JSON(http.StatusForbidden, nil)
	}
	updateUserProfileService.UserID = userID
	if err := c.ShouldBind(&updateUserProfileService); err != nil { // 这里通过shouldBind进行绑定，可以将值传到userRegister变量中
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
	} else {
		res := updateUserProfileService.Update(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// GetUserProfile 获取用户信息
func GetUserProfile(c *gin.Context) { //todo:要不要弄一个同时读user
	var getUserProfileService service.GetUserService
	if err := c.ShouldBind(&getUserProfileService); err != nil { // 通过shouldBind进行绑定
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
	} else {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		}
		uid, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		}
		getUserProfileService.UserID = uint(uid)
		res := getUserProfileService.Get(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}
