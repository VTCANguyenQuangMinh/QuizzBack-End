package controller

import (
	"Quizz/dal"
	"Quizz/model"

	"github.com/gin-gonic/gin"
)

//SetupItemRouter for item router
func SetupItemRouter(r *gin.Engine) {

	//Read
	r.GET("/Quizz", func(c *gin.Context) {
		ListQuizz := make([]model.Quizz, 0)
		listQuestions, err := dal.GetAllQuestions()
		if err != nil {
			c.JSON(500, gin.H{
				"messages": "Get all questions faise!",
			})
		} else {
			for _, question := range listQuestions {
				ListAnswer, _ := dal.GetAnswerByQuestID(question.QuestionID)
				Quizz := model.Quizz{}
				Quizz.Question.QuestionID = question.QuestionID
				Quizz.Question.Question = question.Question
				Quizz.Question.TypeAnswer = question.TypeAnswer

				Quizz.ListAnswer = ListAnswer
				ListQuizz = append(ListQuizz, Quizz)
			}

			c.JSON(200, ListQuizz)
		}
	})

}
