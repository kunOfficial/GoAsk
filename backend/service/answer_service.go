package service

import (
	"GoAsk/cache"
	"GoAsk/dao"
	"GoAsk/serializer"
	e "GoAsk/utils/error"
	"context"
)

type AnswerService struct {
	ID            uint
	AnswererID    uint
	QuestionID    uint   `json:"question_id" form:"question_id"`
	Content       string `json:"content" form:"content" bind:"min=0, max=200"`
	QuestionTitle string
}

type DeleteAnswerService struct {
	ID     uint
	UserID uint
}

type GetAnswersService struct {
	UserID   uint64
	PageSize uint `json:"page_size" form:"page_size"`
	PageNum  uint `json:"page_num" form:"page_num"`
}

type AnswerLikeService struct {
	ID     uint
	UserID uint
}

// GetLikes 查看user是否点赞
func (service *AnswerLikeService) GetLikes(ctx context.Context) serializer.Response {
	// 先查缓存里有没有
	cacheCli := cache.NewCacheClient(ctx)
	foundKey, likes, err := cacheCli.GetLikes(service.ID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorCache, err)
	}
	if !foundKey { // 如果没找到key, 那就去数据库里找找
		dbClient := dao.NewDbClient(ctx)
		uids, err := dbClient.GetAnswerLikes(service.ID)
		if err != nil {
			return serializer.BuildResponse(e.ErrorDataBase, err)
		}
		// 从数据库读取likes
		likes = int64(len(uids))
		// 将数据库信息更新到缓存中
		err = cacheCli.LoadLikes(service.ID, uids)
		if err != nil {
			return serializer.BuildResponse(e.ErrorCache, err)
		}
	}
	// 最后再通过缓存来查 liked 信息
	if service.ID == 0 {
		return serializer.BuildResponse(e.SUCCESS, map[string]interface{}{"is_liked": false, "likes": likes})
	} else {
		_, liked, err := cacheCli.IsLiked(service.ID, service.UserID)
		if err != nil {
			return serializer.BuildResponse(e.ErrorCache, nil)
		}
		return serializer.BuildResponse(e.SUCCESS, map[string]interface{}{"is_liked": liked, "likes": likes})
	}
}

// SetLike 点赞
func (service *AnswerLikeService) SetLike(ctx context.Context) serializer.Response {
	cacheCli := cache.NewCacheClient(ctx)
	liked, err := cacheCli.AddLike(service.ID, service.UserID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorCache, nil)
	}
	if liked { // 如果已经点过赞
		return serializer.BuildResponse(e.ErrorAlreadyLiked, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, nil)
}

// CancelLike 取消点赞
func (service *AnswerLikeService) CancelLike(ctx context.Context) serializer.Response {
	cacheCli := cache.NewCacheClient(ctx)
	inAddSet, err := cacheCli.RemoveLike(service.ID, service.UserID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorCache, err)
	}
	if !inAddSet { // 如果在缓存中，那就需要将其从数据库中删除
		dbClient := dao.NewDbClient(ctx)
		err := dbClient.DeleteLike(service.ID, service.UserID)
		if err != nil {
			return serializer.BuildResponse(e.ErrorDataBase, err)
		}
	}
	return serializer.BuildResponse(e.SUCCESS, nil)
}

// Post 发布回答
func (service *AnswerService) Post(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	user, _, err := cli.GetUserByUserID(service.AnswererID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	question, found, err := cli.GetQuestionByID(service.QuestionID, false)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false { // 问题不存在
		return serializer.BuildResponse(e.ErrorQuestionNotExist, nil)
	}
	answer, err := cli.CreateAnswer(user.ID, service.QuestionID, service.Content, question.Title, user.NickName)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildAnswer(answer))
}

// Get 获取单个 answer
func (service *AnswerService) Get(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	q, found, err := cli.GetAnswerByID(service.ID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false {
		return serializer.BuildResponse(e.ErrorAnswerNotExist, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildAnswer(q))
}

// Get 获取所有 answer
func (service *GetAnswersService) Get(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	as, err := cli.GetAnswers(service.PageNum, service.PageSize)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildAnswers(as))
}

// GetByUserID 获取所有 answer
func (service *GetAnswersService) GetByUserID(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	as, err := cli.GetAnswersByUserID(service.PageNum, service.PageSize, uint(service.UserID))
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildAnswers(as))
}

// Delete 删除问题
func (service *DeleteAnswerService) Delete(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	answer, found, err := cli.GetAnswerByID(service.ID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false {
		return serializer.BuildResponse(e.ErrorAnswerNotExist, nil)
	}
	if answer.AnswererID != service.UserID { // 判断用户的身份是否与 AnswererID 吻合
		return serializer.BuildResponse(e.AccessDenied, nil)
	}
	found, err = cli.DeleteAnswer(service.ID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false { // 其实可以不检查found的
		return serializer.BuildResponse(e.ErrorQuestionNotExist, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, nil)
}

// Update 修改回答
func (service *AnswerService) Update(ctx context.Context) serializer.Response {
	cli := dao.NewDbClient(ctx)
	answer, found, err := cli.GetAnswerByID(service.ID)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	if found == false {
		return serializer.BuildResponse(e.ErrorAnswerNotExist, nil)
	}
	if answer.AnswererID != service.AnswererID { // 判断用户的身份是否与 AnswererID 吻合
		return serializer.BuildResponse(e.AccessDenied, nil)
	}
	newAnswer, err := cli.UpdateAnswer(service.ID, service.Content)
	if err != nil {
		return serializer.BuildResponse(e.ErrorDataBase, nil)
	}
	return serializer.BuildResponse(e.SUCCESS, serializer.BuildAnswer(newAnswer))
}
