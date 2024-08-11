package models

import "github.com/google/uuid"

// Question represents a question in a quiz
type Question struct {
	ID       uuid.UUID `json:"id"`
	QuizID   uuid.UUID `json:"quiz_id"`
	Type     string    `json:"type"` // "mcq", "short", "true_false"
	Question string    `json:"question"`
	Options  []string  `json:"options,omitempty"` // Only for MCQ
	Answer   string    `json:"answer"`
}
