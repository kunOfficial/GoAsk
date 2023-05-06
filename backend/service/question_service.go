package service

import (
	"GoAsk/cache"
	"GoAsk/dao"
	"GoAsk/serializer"
	e "GoAsk/utils/error"
	"context"
)

type QuestionService struct {
	ID           uint
	QuestionerID uint
	Title        string `json:"title" form:"title" bind:"min=5,max=20"`
	Description  string `json:"description" form:"description" bind:"min=0, max=200"`
}

type GetQuestionService struct {
	ID    uint
	Brief bool
}

type GetQuestionsService struct {
	PageSize uint `json:"page_size" form:"page_size"`
	PageNum  uint `json:"page_num" form:"page_num"`
	UserID   uint64
	Filter   string
}

type DeleteQuestionService struct {
	UserID uint
	ID     uint
}

func (service *QuestionService) Post(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	user, _, err := cli.GetUserByUserID(service.QuestionerID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	question, err := cli.CreateQuestion(user.ID, service.Title, service.Description, user.NickName)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildQuestion(question))
}

// Update 几乎和 Post 一样
func (service *QuestionService) Update(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	question, found, err := cli.GetQuestionByID(service.ID, true)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false {
		return serializer.BuildResponse(e.ErrorQuestionNotExist, nil)
	}
	if question.QuestionerID != service.QuestionerID { // 判断用户的身份是否与 QuestionerID 吻合
		return serializer.BuildResponse(e.AccessDenied, nil)
	}
	newQuestion, err := cli.UpdateQuestion(service.ID, service.Title, service.Description)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildQuestion(newQuestion))
}

// Get 获取单个question
func (service *GetQuestionService) Get(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	q, found, err := cli.GetQuestionByID(service.ID, service.Brief)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false {
		return serializer.BuildResponse(e.ErrorQuestionNotExist, nil)
	}
	cacheCli := cache.NewCacheClient(ctx)
	if service.Brief == false { // 如果是用户进入了界面
		err = cacheCli.AddQuestionView(int(q.ID))
		if err != nil {
			return serializer.BuildResponse(e.ErrorCache, nil)
		}
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildQuestion(q))
}

// GetPage 获取所有question
func (service *GetQuestionsService) GetPage(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	qs, err := cli.GetQuestions(service.PageNum, service.PageSize)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildBriefQuestions(qs))
}

// Search 根据title和description搜索question
func (service *GetQuestionsService) Search(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	qs, err := cli.GetQuestionsWithCond(service.PageNum, service.PageSize, service.Filter)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildBriefQuestions(qs))
}

// GetByUserID 获取某个用户的所有问题
func (service *GetQuestionsService) GetByUserID(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	qs, err := cli.GetQuestionsByUserID(service.PageNum, service.PageSize, uint(service.UserID))
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildBriefQuestions(qs))
}

func (service *DeleteQuestionService) Delete(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	question, found, err := cli.GetQuestionByID(service.ID, true)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false {
		return serializer.BuildResponse(e.ErrorAnswerNotExist, nil)
	}
	if question.QuestionerID != service.UserID { // 判断用户的身份是否与 QuestionerID 吻合
		return serializer.BuildResponse(e.AccessDenied, nil)
	}
	found, err = cli.DeleteQuestion(service.ID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false { // 其实可以不检查found的
		return serializer.BuildResponse(e.ErrorQuestionNotExist, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, nil)
}
