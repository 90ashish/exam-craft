package models

import "github.com/google/uuid"

// Quiz represents a quiz in the system
type Quiz struct {
	ID           uuid.UUID  `json:"id"`
	Title        string     `json:"title"`
	InstructorID uuid.UUID  `json:"instructor_id"`
	Questions    []Question `json:"questions"`
}
