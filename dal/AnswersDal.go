package dal

import (
	"Quizz/model"
)

// GetAnswerByQuestID Get list answers of question
func GetAnswerByQuestID(QuestID int64) ([]model.Answer, error) {
	GetConnection()
	stmt, _ := db.Prepare("select answer_id, answer from answers where question_id=?;")

	defer CloseStmt(stmt)

	rows, err := stmt.Query(QuestID)
	listAnswers := make([]model.Answer, 0)
	for rows.Next() {
		var Answer model.Answer
		err := rows.Scan(&Answer.AnswerID, &Answer.Answer)
		if err != nil {
			return listAnswers, err
		}

		listAnswers = append(listAnswers, Answer)

	}
	return listAnswers, err
}

// GetCorrectAnswer  Get list correct answers of question
func GetCorrectAnswer(QuestID int64) ([]model.CorrectAnswers, error) {
	GetConnection()
	stmt, _ := db.Prepare("select answer_id, answer, score from answers where question_id=? AND correct=1;")

	defer CloseStmt(stmt)

	rows, err := stmt.Query(QuestID)
	listCorrectAnswers := make([]model.CorrectAnswers, 0)
	if err != nil {
		return listCorrectAnswers, err
	}

	for rows.Next() {
		var CorrectAnswer model.CorrectAnswers
		err := rows.Scan(&CorrectAnswer.AnswerID, &CorrectAnswer.Answer, &CorrectAnswer.Score)
		if err != nil {
			return listCorrectAnswers, err
		}

		listCorrectAnswers = append(listCorrectAnswers, CorrectAnswer)
	}
	return listCorrectAnswers, err
}

// GetAnswerByAnswerID return model.Answer data
func GetAnswerByAnswerID(AnswerID float64) (model.Answer, error) {
	GetConnection()
	stmt, err := db.Prepare("select answer_id, answer from answers where answer_id=?;")
	defer stmt.Close()
	var Answer model.Answer
	stmt.QueryRow(AnswerID).Scan(&Answer.AnswerID, &Answer.Answer)
	return Answer, err
}
