package dal

import (
	"Quizz/model"
)

// GetAnswerByQuestID Get list answer of question
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
