package controller

import (
	"Quizz/dal"
	"Quizz/model"
)

// ConvertListIDAnswersToListAnswers return list model.answer form
func ConvertListIDAnswersToListAnswers(arr []float64) []model.Answer {
	listAnswers := make([]model.Answer, 0)
	for _, answerID := range arr {
		var answer model.Answer
		answer, _ = dal.GetAnswerByAnswerID(answerID)
		listAnswers = append(listAnswers, answer)
	}
	return listAnswers
}
