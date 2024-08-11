package repositories

import (
	"context"
	"exam-craft/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// QuestionRepository handles database operations for questions
type QuestionRepository struct {
	db *sqlx.DB
}

// NewQuestionRepository creates a new QuestionRepository
func NewQuestionRepository(db *sqlx.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

// CreateQuestion inserts a new question into the database
func (r *QuestionRepository) CreateQuestion(ctx context.Context, question *models.Question) error {
	question.ID = uuid.New()
	query := `INSERT INTO questions (id, quiz_id, type, question, options, answer) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, question.ID, question.QuizID, question.Type, question.Question, pq.Array(question.Options), question.Answer)
	return err
}

// GetQuestionsByQuizID retrieves questions by quiz ID from the database
func (r *QuestionRepository) GetQuestionsByQuizID(ctx context.Context, quizID uuid.UUID) ([]models.Question, error) {
	var questions []models.Question
	query := `SELECT * FROM questions WHERE quiz_id = $1`
	err := r.db.SelectContext(ctx, &questions, query, quizID)
	return questions, err
}
