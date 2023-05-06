package serializer

import (
	"GoAsk/model"
)

type SerializedBriefQuestion struct {
	ID                 uint   `json:"id"`
	QuestionerID       uint   `json:"questioner_id"`
	QuestionerNickName string `json:"questioner_nick_name"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	ViewNum            uint   `json:"view"`
	UpdatedAt          string `json:"updated_at"`
}

type SerializedQuestion struct {
	SerializedBriefQuestion
	Answers []SerializedAnswer `json:"answers"`
}

func BuildQuestion(question model.Question) SerializedQuestion {
	return SerializedQuestion{
		SerializedBriefQuestion: BuildBriefQuestion(question),
		Answers:                 BuildAnswers(question.Answers),
	}
}

func BuildBriefQuestion(question model.Question) SerializedBriefQuestion {
	return SerializedBriefQuestion{
		ID:                 question.ID,
		QuestionerID:       question.QuestionerID,
		QuestionerNickName: question.QuestionerNickName,
		Description:        question.Description,
		Title:              question.Title,
		ViewNum:            question.ViewNum,
		UpdatedAt:          question.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func BuildBriefQuestions(questions []model.Question) []SerializedBriefQuestion {
	result := make([]SerializedBriefQuestion, len(questions))
	for index, question := range questions {
		result[index] = BuildBriefQuestion(question)
	}
	return result
}
