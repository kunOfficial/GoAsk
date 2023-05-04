package v1

import (
	"GoAsk/serializer"
	"GoAsk/service"
	e "GoAsk/utils/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetSpecificQuestion 获得某个问题
func GetSpecificQuestion(c *gin.Context) {
	var getQuestionByIDService service.GetQuestionService
	if err := c.ShouldBind(&getQuestionByIDService); err != nil { // 通过shouldBind进行绑定
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
	} else {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		}
		qid, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		}
		getQuestionByIDService.ID = uint(qid)
		// 通过 brief 的 Query 参数来决定是否查 answers
		brief := c.Query("brief")
		if brief == "" || brief == "false" || brief == "true" {
			if brief == "true" {
				getQuestionByIDService.Brief = true
			} else { // 默认 brief
				getQuestionByIDService.Brief = false
			}
		} else {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		}
		res := getQuestionByIDService.Get(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// PostQuestion 发布一个问题
func PostQuestion(c *gin.Context) {
	var postQuestionService service.QuestionService
	if err := c.ShouldBind(&postQuestionService); err != nil { // 通过shouldBind进行绑定
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
	} else {
		postQuestionService.QuestionerID = c.MustGet("user_id").(uint)
		res := postQuestionService.Post(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// GetQuestions 获取所有/某个用户的问题
func GetQuestions(c *gin.Context) {
	var getQuestionsService service.GetQuestionsService
	pageNum, err1 := strconv.Atoi(c.Query("page_num"))
	pageSize, err2 := strconv.Atoi(c.Query("page_size"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
	} else {
		getQuestionsService.PageNum = uint(pageNum)
		getQuestionsService.PageSize = uint(pageSize)
		if c.Query("uid") != "" { //
			var err3 error
			getQuestionsService.UserID, err3 = strconv.ParseUint(c.Query("uid"), 10, 32)
			if err3 != nil {
				c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
				return
			}
			res := getQuestionsService.GetByUserID(c.Request.Context())
			c.JSON(http.StatusOK, res)
		} else {
			res := getQuestionsService.GetPage(c.Request.Context())
			c.JSON(http.StatusOK, res)
		}
	}
}

// SearchQuestions 根据filter来搜索问题
func SearchQuestions(c *gin.Context) {
	var searchQuestionsService service.GetQuestionsService
	filter := c.Query("filter")
	if filter == "" {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
		return
	}
	pageNum, err1 := strconv.Atoi(c.Query("page_num"))
	pageSize, err2 := strconv.Atoi(c.Query("page_size"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, nil))
		return
	}
	searchQuestionsService = service.GetQuestionsService{
		PageSize: uint(pageSize),
		PageNum:  uint(pageNum),
		Filter:   filter,
	}
	res := searchQuestionsService.Search(c.Request.Context())
	c.JSON(http.StatusOK, res)
}

// UpdateQuestion 更新问题
func UpdateQuestion(c *gin.Context) {
	var updateQuestionService service.QuestionService
	userID := c.MustGet("user_id").(uint)
	updateQuestionService.QuestionerID = userID
	if err := c.ShouldBind(&updateQuestionService); err != nil { // 通过shouldBind进行绑定
		c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
		return
	} else {
		if qid, err := strconv.Atoi(c.Param("id")); err != nil {
			c.JSON(http.StatusBadRequest, serializer.BuildResponse(e.InvalidParams, err))
			return
		} else {
			updateQuestionService.ID = uint(qid)
		}
		res := updateQuestionService.Update(c.Request.Context())
		c.JSON(http.StatusOK, res)
	}
}

// DeleteQuestion 删除问题
func DeleteQuestion(c *gin.Context) {
	var deleteService service.DeleteQuestionService
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
