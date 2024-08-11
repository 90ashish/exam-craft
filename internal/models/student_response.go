package models

import "github.com/google/uuid"

// StudentResponse represents a student's response to a quiz
type StudentResponse struct {
	ID          uuid.UUID  `json:"id"`
	StudentID   uuid.UUID  `json:"student_id"`
	QuizID      uuid.UUID  `json:"quiz_id"`
	Responses   []Response `json:"responses"`
	SubmittedAt int64      `json:"submitted_at"`
}

// Response represents an individual response to a question
type Response struct {
	QuestionID uuid.UUID `json:"question_id"`
	Answer     string    `json:"answer"`
}
