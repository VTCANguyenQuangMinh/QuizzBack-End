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

type UsersAnswers struct {
	QuestionID int64
	AnswersID  []float64
}

// Question is bla blabla...
type Question struct {
	QuestionID int64  `json:"questionId"`
	Question   string `json:"question"`
	TypeAnswer string `json:"type"`
}

// Result is bla blo ble...
type Result struct {
	QuestionID        int64            `json:"questionId"`
	ListCorrectAnswer []CorrectAnswers `json:"CorrectAnswers"`
	ListUsersAnswer   []Answer         `json:"UsersAnswers"`
}

// CorrectAnswers hdfvk
type CorrectAnswers struct {
	AnswerID float64 `json:"answerId"`
	Answer   string  `json:"answer"`
	Score    float32 `json:"score"`
}
