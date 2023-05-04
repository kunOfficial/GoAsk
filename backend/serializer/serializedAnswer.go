package serializer

import "GoAsk/model"

type SerializedAnswer struct {
	ID               uint   `json:"id"`
	QuestionID       uint   `json:"question_id"`
	QuestionTitle    string `json:"question_title"`
	AnswererNickName string `json:"answerer_nick_name"`
	AnswererID       uint   `json:"answerer_id"`
	Content          string `json:"content"`
	UpdatedAt        string `json:"updated_at"`
}

func BuildAnswer(answer model.Answer) SerializedAnswer {
	return SerializedAnswer{
		ID:               answer.ID,
		QuestionID:       answer.QuestionID,
		QuestionTitle:    answer.QuestionTitle,
		AnswererID:       answer.AnswererID,
		AnswererNickName: answer.AnswererNickName,
		Content:          answer.Content,
		UpdatedAt:        answer.UpdatedAt.Format("2006-01-02 15:04:05"),
		// 默认false
	}
}

func BuildAnswers(answers []model.Answer) []SerializedAnswer {
	result := make([]SerializedAnswer, len(answers))
	for index, answer := range answers {
		result[index] = BuildAnswer(answer)
	}
	return result
}
