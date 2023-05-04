package v1

import (
	"GoAsk/serializer"
	"GoAsk/service"
	e "GoAsk/utils/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetSpecificAnswer 获得某个回答
func GetSpecificAnswer(c *gin.Context) {
	var getAnswerByIDService service.AnswerService
	if err := c.ShouldBind(&getAnswerByIDService); err != nil { // 通过shouldBind进行绑定
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
	} else {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
			return
		}
		qid, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
			return
		}
		getAnswerByIDService.ID = uint(qid)
		res := getAnswerByIDService.Get(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// PostAnswer 发布一个回答
func PostAnswer(c *gin.Context) {
	var postAnswerService service.AnswerService
	if err := c.ShouldBind(&postAnswerService); err != nil { // 通过shouldBind进行绑定
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
	} else {
		postAnswerService.AnswererID = c.MustGet("user_id").(uint)
		res := postAnswerService.Post(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// GetAnswers 获得一页回答
func GetAnswers(c *gin.Context) {
	var getAnswersService service.GetAnswersService
	pageNum, err1 := strconv.Atoi(c.Query("page_num"))
	pageSize, err2 := strconv.Atoi(c.Query("page_size"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
	} else {
		getAnswersService.PageNum = uint(pageNum)
		getAnswersService.PageSize = uint(pageSize)
		if c.Query("uid") != "" {
			var err3 error
			getAnswersService.UserID, err3 = strconv.ParseUint(c.Query("uid"), 10, 32)
			if err3 != nil {
				c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
				return
			}
			res := getAnswersService.GetByUserID(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			res := getAnswersService.Get(c.Request.Context())
			c.JSON(http.StatusOK, res)
		}
	}
}

// GetAnswerLikes 获取
func GetAnswerLikes(c *gin.Context) {
	var getAnswerLikeService service.AnswerLikeService
	answerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
		return
	}
	if c.Query("uid") == "" { // 如果 uid 为空，说明为游客(id=0)
		getAnswerLikeService.UserID = 0
	} else {
		userID, err := strconv.ParseUint(c.Query("uid"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
			return
		}
		getAnswerLikeService.UserID = uint(userID)
	}
	getAnswerLikeService.ID = uint(answerID)
	res := getAnswerLikeService.GetLikes(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

// LikeAnswer 点赞
func LikeAnswer(c *gin.Context) {
	answerID, err1 := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, err2 := strconv.ParseUint(c.Query("uid"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
		return
	}
	var likeAnswerService = service.AnswerLikeService{
		ID:     uint(answerID),
		UserID: uint(userID),
	}
	res := likeAnswerService.SetLike(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

func CancelLikeAnswer(c *gin.Context) {
	answerID, err1 := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, err2 := strconv.ParseUint(c.Query("uid"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
		return
	}
	var cancelLikeService = service.AnswerLikeService{
		ID:     uint(answerID),
		UserID: uint(userID),
	}
	res := cancelLikeService.CancelLike(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

// UpdateAnswer 修改回答
func UpdateAnswer(c *gin.Context) {
	var updateService service.AnswerService
	userID := c.MustGet("user_id").(uint)
	updateService.AnswererID = userID
	if err := c.ShouldBind(&updateService); err != nil { // 通过shouldBind进行绑定
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		return
	} else {
		if qid, err := strconv.Atoi(c.Param("id")); err != nil {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
			return
		} else {
			updateService.ID = uint(qid)
		}
		res := updateService.Update(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// DeleteAnswer 删除回答
func DeleteAnswer(c *gin.Context) {
	var deleteService service.DeleteAnswerService
	if qid, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		return
	} else {
		deleteService.ID = uint(qid)
	}
	userID := c.MustGet("user_id").(uint)
	deleteService.UserID = userID
	res := deleteService.Delete(c.Request.Context())
	c.JSON(http.StatusOK, res)
}
