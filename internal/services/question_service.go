package services

import (
	"context"
	"exam-craft/internal/models"
	"exam-craft/internal/repositories"

	"github.com/google/uuid"
)

// QuestionService handles business logic for questions
type QuestionService struct {
	questionRepo *repositories.QuestionRepository
}

// NewQuestionService creates a new QuestionService
func NewQuestionService(questionRepo *repositories.QuestionRepository) *QuestionService {
	return &QuestionService{questionRepo: questionRepo}
}

// CreateQuestion creates a new question
func (s *QuestionService) CreateQuestion(ctx context.Context, question *models.Question) error {
	return s.questionRepo.CreateQuestion(ctx, question)
}

// GetQuestionsByQuizID retrieves questions by quiz ID
func (s *QuestionService) GetQuestionsByQuizID(ctx context.Context, quizID uuid.UUID) ([]models.Question, error) {
	return s.questionRepo.GetQuestionsByQuizID(ctx, quizID)
}
