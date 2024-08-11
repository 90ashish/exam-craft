package services

import (
	"context"
	"exam-craft/internal/models"
	"exam-craft/internal/repositories"

	"github.com/google/uuid"
)

// StudentService handles business logic for students
type StudentService struct {
	studentResponseRepo *repositories.StudentResponseRepository
	quizRepo            *repositories.QuizRepository
}

// NewStudentService creates a new StudentService
func NewStudentService(studentResponseRepo *repositories.StudentResponseRepository, quizRepo *repositories.QuizRepository) *StudentService {
	return &StudentService{studentResponseRepo: studentResponseRepo, quizRepo: quizRepo}
}

// ListAvailableQuizzes lists all available quizzes
func (s *StudentService) ListAvailableQuizzes(ctx context.Context) ([]models.Quiz, error) {
	return s.quizRepo.ListQuizzes(ctx)
}

// SubmitResponse submits a student's response to a quiz
func (s *StudentService) SubmitResponse(ctx context.Context, response *models.StudentResponse) error {
	return s.studentResponseRepo.CreateStudentResponse(ctx, response)
}

// GetResponsesByQuizID retrieves student responses by quiz ID
func (s *StudentService) GetResponsesByQuizID(ctx context.Context, quizID uuid.UUID) ([]models.StudentResponse, error) {
	return s.studentResponseRepo.GetResponsesByQuizID(ctx, quizID)
}
