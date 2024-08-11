package utils

import (
	"errors"
	"regexp"
)

// Regular expression for email validation
var emailRegex = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)

// ValidateEmail validates the email format
func ValidateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword validates the password length
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

// ValidateUserRole validates the user role
func ValidateUserRole(role string) error {
	if role != "student" && role != "instructor" {
		return errors.New("invalid user role")
	}
	return nil
}

// ValidateQuizTitle validates the quiz title
func ValidateQuizTitle(title string) error {
	if len(title) == 0 {
		return errors.New("quiz title cannot be empty")
	}
	return nil
}

// ValidateQuestionType validates the question type
func ValidateQuestionType(questionType string) error {
	validTypes := map[string]bool{
		"mcq":        true,
		"short":      true,
		"true_false": true,
	}
	if !validTypes[questionType] {
		return errors.New("invalid question type")
	}
	return nil
}
