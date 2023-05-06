package test

import (
	"GoAsk/serializer"
	"GoAsk/service"
	e "GoAsk/utils/error"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateAndGetQuestionService(t *testing.T) {
	InitForTest()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	create_question_service := service.QuestionService{
		ID:           0,
		QuestionerID: 10,
		Title:        "a title",
		Description:  "no description",
	}
	response := create_question_service.Post(ctx)
	sq := response.Data.(serializer.SerializedQuestion)
	assert.Equal(t, sq.Title, "a title")
	assert.Equal(t, sq.Description, "no description")
	assert.Equal(t, sq.QuestionerID, 10)
}

func TestUpdateQuestionService(t *testing.T) {
	InitForTest()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//create_user_serive := service.UserService{
	//	//	UserName: "test",
	//	//	Password: "mekk",
	//	//	NickName: "mekk",
	//	//}
	//	//response := create_user_serive.Register(ctx)
	//	//assert.Equal(t, response.Status, e.SUCCESS)
	get_user_service := service.GetUserService{UserID: 1}
	response := get_user_service.Get(ctx)
	user_id := response.Data.(serializer.SerializedUser).ID

	create_question_service := service.QuestionService{
		ID:           0,
		QuestionerID: user_id,
		Title:        "a title",
		Description:  "no description",
	}
	response = create_question_service.Post(ctx)
	sq := response.Data.(serializer.SerializedQuestion)
	q_id := sq.ID
	assert.Equal(t, sq.QuestionerID, uint(1))

	t.Run("update-question-with-wrong-questionerID", func(t *testing.T) {
		update_question_service := service.QuestionService{
			ID:           q_id,
			QuestionerID: q_id + 1,
			Title:        "aa",
			Description:  "bb",
		}
		response := update_question_service.Update(ctx)
		assert.Equal(t, response.Status, e.AccessDenied)
	})

	t.Run("update-question", func(t *testing.T) {
		update_question_service := service.QuestionService{
			ID:           q_id,
			QuestionerID: 1,
			Title:        "Test5",
			Description:  "Test6",
		}
		response := update_question_service.Update(ctx)
		assert.Equal(t, response.Status, e.SUCCESS)
		sq := response.Data.(serializer.SerializedQuestion)
		assert.Equal(t, sq.ID, q_id)
		assert.Equal(t, sq.QuestionerID, user_id)
		assert.Equal(t, sq.Title, "Test5")
		assert.Equal(t, sq.Description, "Test6")
	})

	delete_question_service := service.DeleteQuestionService{
		UserID: user_id,
		ID:     q_id,
	}
	response = delete_question_service.Delete(ctx)
	assert.Equal(t, response.Status, e.SUCCESS)

}
