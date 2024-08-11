package models

import "github.com/google/uuid"

// User represents a user in the system
type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"` // Hashed password
	Role     string    `json:"role"`     // "student" or "instructor"
}
