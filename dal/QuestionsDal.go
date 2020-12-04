package dal

import (
	"Quizz/model"
)

// GetAllQuestions get all questions from Quizz
func GetAllQuestions() ([]model.Question, error) {
	GetConnection()
	stmt, _ := db.Prepare("select * from questions;")

	defer stmt.Close()

	rows, err := stmt.Query()
	listQuestions := make([]model.Question, 0)
	for rows.Next() {
		var Question model.Question
		err := rows.Scan(&Question.QuestionID, &Question.Question, &Question.TypeAnswer)
		if err != nil {
			return listQuestions, err
		}

		listQuestions = append(listQuestions, Question)

	}
	return listQuestions, err
}

// GetAllIDQuestions get all ID questions from Quizz
func GetAllIDQuestions() ([]int64, error) {
	GetConnection()
	stmt, _ := db.Prepare("select question_id from questions;")

	defer stmt.Close()

	rows, err := stmt.Query()
	listIDQuestions := make([]int64, 0)
	for rows.Next() {
		var IDQuestion int64
		err := rows.Scan(&IDQuestion)
		if err != nil {
			return listIDQuestions, err
		}

		listIDQuestions = append(listIDQuestions, IDQuestion)

	}
	return listIDQuestions, err
}
