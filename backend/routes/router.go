package routes

import (
	v1 "GoAsk/api/v1"
	"GoAsk/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))

	//r.StaticFS("/avatar", http.Dir("./static/avatar"))
	r.Static("/avatar", ".\\static\\avatar")
	r.Use(sessions.Sessions("mySession", store), middleware.Cors())
	apiV1 := r.Group("api/v1")
	{
		apiV1.POST("/user/register", v1.UserRegister)
		apiV1.POST("/user/login", v1.UserLogin)
		apiV1.GET("/users/:id", v1.GetUserProfile)

		apiV1.GET("/questions/:id", v1.GetSpecificQuestion)
		apiV1.GET("/questions", v1.GetQuestions)

		apiV1.GET("/answers/:id", v1.GetSpecificAnswer)
		apiV1.GET("/answers", v1.GetAnswers)

		apiV1.GET("/avatars/:id", v1.GetAvatar)

		apiV1.GET("/answers/:id/likes", v1.GetAnswerLikes)

		// TODO: 暂时只支持搜索问题
		// 其实要完全按照 RestFul 规范的话，这里应该集成到 /questions 中，不过我感觉没必要，这样做会复杂
		apiV1.GET("/questions/search", v1.SearchQuestions)

		authed := apiV1.Group("/")
		authed.Use(middleware.JwtAuth())
		{
			// 用户操作
			authed.PATCH("/users/:id", v1.UpdateUserProfile)
			authed.POST("/users/")

			// 问题操作
			authed.POST("/questions", v1.PostQuestion)
			authed.PATCH("/questions/:id", v1.UpdateQuestion)
			authed.DELETE("/questions/:id", v1.DeleteQuestion)

			// 回答操作
			authed.POST("/answers", v1.PostAnswer)
			authed.PATCH("/answers/:id", v1.UpdateAnswer)
			authed.DELETE("/answers/:id", v1.DeleteAnswer)
			apiV1.POST("/answers/:id/likes", v1.LikeAnswer)
			apiV1.DELETE("/answers/:id/likes", v1.CancelLikeAnswer)

			// 上传头像
			authed.POST("/avatars", v1.UploadAvatar)
		}
	}
	return r
}
