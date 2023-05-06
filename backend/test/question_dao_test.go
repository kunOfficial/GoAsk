package test

import (
	"GoAsk/dao"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func InitForTest() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/GoAsk_test_db?charset=utf8mb4&parseTime=true"
	db := dao.Connect(dsn)
	dao.Init(db)
	//dao.AutoMigration(db)
}

func TestCreateAndGetQuestion(t *testing.T) {
	InitForTest()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client := dao.NewDbClient(ctx)

	new_question, err := client.CreateQuestion(1, "this is question title", "hello", "mekk")
	assert.Nil(t, err)
	question, found, err := client.GetQuestionByID(1, false)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.Equal(t, new_question.Title, question.Title)
	assert.Equal(t, new_question.QuestionerID, question.QuestionerID)
	assert.Equal(t, new_question.Description, question.Description)
	assert.Equal(t, new_question.QuestionerNickName, question.QuestionerNickName)
}

func TestUpdateQuestion(t *testing.T) {
	InitForTest()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client := dao.NewDbClient(ctx)

	orignal_question, err := client.CreateQuestion(0, "Original Title", "Original Description", "mekk")
	assert.Nil(t, err)

	t.Run("update-title-and-description", func(t *testing.T) {
		q1, err := client.UpdateQuestion(orignal_question.ID, "Test1", "Test2")
		assert.Nil(t, err)
		assert.Equal(t, q1.Title, "Test1")
		assert.Equal(t, q1.Description, "Test2")
	})

	t.Run("update-only-title", func(t *testing.T) {
		q2, err := client.UpdateQuestion(orignal_question.ID, "Test3", "Test2")
		assert.Nil(t, err)
		assert.Equal(t, q2.Title, "Test3")
		assert.Equal(t, q2.Description, "Test2")
	})

	t.Run("update-only-description", func(t *testing.T) {
		q3, err := client.UpdateQuestion(orignal_question.ID, "", "Test4")
		assert.Nil(t, err)
		assert.Equal(t, q3.Title, "Test3")
		assert.Equal(t, q3.Description, "Test4")
	})

	found, err := client.DeleteQuestion(orignal_question.ID)
	assert.True(t, found)
	assert.Nil(t, err)
}
