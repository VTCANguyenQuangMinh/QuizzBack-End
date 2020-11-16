package model

// Quizz is bla blabla...
type Quizz struct {
	Question   Question `json:"question"`
	ListAnswer []Answer `json:"answers"`
}

// Answer is bla blabla...
type Answer struct {
	AnswerID float64 `json:"answerId"`
	Answer   string  `json:"answer"`
}

// Question is bla blabla...
type Question struct {
	QuestionID int64  `json:"questionId"`
	Question   string `json:"question"`
	TypeAnswer string `json:"type"`
}
