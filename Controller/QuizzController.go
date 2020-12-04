package controller

import (
	"Quizz/dal"
	"Quizz/model"
	"fmt"

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

	r.POST("/Quizz", func(c *gin.Context) {

		ListResult := make([]model.Result, 0)
		ListUsersAnswers := make([]model.UsersAnswers, 0)

		if err := c.ShouldBindJSON(&ListUsersAnswers); err == nil {
			ListIDQuestions, err := dal.GetAllIDQuestions()
			if err != nil {
				c.JSON(500, gin.H{
					"messages": "Get all correct answers error",
				})
			} else {
				for _, IDQuest := range ListIDQuestions {
					var Result model.Result

					for _, UsersAnswer := range ListUsersAnswers {
						if IDQuest == UsersAnswer.QuestionID {
							Result.ListUsersAnswer = ConvertListIDAnswersToListAnswers(UsersAnswer.AnswersID)
						}
					}
					Result.ListCorrectAnswer, err = dal.GetCorrectAnswer(IDQuest)
					fmt.Println(err)
					Result.QuestionID = IDQuest
					ListResult = append(ListResult, Result)
				}

				c.JSON(200, ListResult)

			}

		}
	})

}
