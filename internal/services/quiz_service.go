package services

import (
	"context"
	"exam-craft/internal/models"
	"exam-craft/internal/repositories"

	"github.com/google/uuid"
)

// QuizService handles business logic for quizzes
type QuizService struct {
	quizRepo *repositories.QuizRepository
}

// NewQuizService creates a new QuizService
func NewQuizService(quizRepo *repositories.QuizRepository) *QuizService {
	return &QuizService{quizRepo: quizRepo}
}

// CreateQuiz creates a new quiz
func (s *QuizService) CreateQuiz(ctx context.Context, quiz *models.Quiz) error {
	quiz.ID = uuid.New()
	return s.quizRepo.CreateQuiz(ctx, quiz)
}

// GetQuiz retrieves a quiz by ID
func (s *QuizService) GetQuiz(ctx context.Context, quizID uuid.UUID) (*models.Quiz, error) {
	return s.quizRepo.GetQuiz(ctx, quizID)
}

// ListQuizzes lists all quizzes
func (s *QuizService) ListQuizzes(ctx context.Context) ([]models.Quiz, error) {
	return s.quizRepo.ListQuizzes(ctx)
}
