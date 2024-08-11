package repositories

import (
	"context"
	"encoding/json"
	"exam-craft/internal/models"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// QuizRepository handles database operations for quizzes
type QuizRepository struct {
	db    *sqlx.DB
	cache *redis.Client
}

// NewQuizRepository creates a new QuizRepository
func NewQuizRepository(db *sqlx.DB, cache *redis.Client) *QuizRepository {
	return &QuizRepository{db: db, cache: cache}
}

// CreateQuiz inserts a new quiz into the database
func (r *QuizRepository) CreateQuiz(ctx context.Context, quiz *models.Quiz) error {
	query := `INSERT INTO quizzes (id, title, instructor_id) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, quiz.ID, quiz.Title, quiz.InstructorID)
	if err != nil {
		return err
	}
	r.updateCache(ctx, quiz.ID)
	return nil
}

// GetQuiz retrieves a quiz by ID from the database or cache
func (r *QuizRepository) GetQuiz(ctx context.Context, quizID uuid.UUID) (*models.Quiz, error) {
	// Try to get from cache
	quizJSON, err := r.cache.Get(ctx, quizID.String()).Result()
	if err == redis.Nil {
		// Not in cache, get from database
		var quiz models.Quiz
		query := `SELECT * FROM quizzes WHERE id = $1`
		err := r.db.GetContext(ctx, &quiz, query, quizID)
		if err != nil {
			return nil, err
		}

		query = `SELECT * FROM questions WHERE quiz_id = $1`
		err = r.db.SelectContext(ctx, &quiz.Questions, query, quizID)
		if err != nil {
			return nil, err
		}

		// Cache the result
		quizBytes, err := json.Marshal(quiz)
		if err != nil {
			return nil, err
		}
		r.cache.Set(ctx, quizID.String(), string(quizBytes), 0)

		return &quiz, nil
	} else if err != nil {
		return nil, err
	}

	// Unmarshal the quiz from cache
	var quiz models.Quiz
	if err := json.Unmarshal([]byte(quizJSON), &quiz); err != nil {
		return nil, err
	}

	return &quiz, nil
}

// ListQuizzes retrieves all quizzes from the database
func (r *QuizRepository) ListQuizzes(ctx context.Context) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	query := `SELECT * FROM quizzes`
	err := r.db.SelectContext(ctx, &quizzes, query)
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

// updateCache updates the cache with the latest quiz data
func (r *QuizRepository) updateCache(ctx context.Context, quizID uuid.UUID) {
	var quiz models.Quiz
	query := `SELECT * FROM quizzes WHERE id = $1`
	err := r.db.GetContext(ctx, &quiz, query, quizID)
	if err != nil {
		log.Printf("Failed to update cache for quiz %s: %v", quizID, err)
		return
	}

	query = `SELECT * FROM questions WHERE quiz_id = $1`
	err = r.db.SelectContext(ctx, &quiz.Questions, query, quizID)
	if err != nil {
		log.Printf("Failed to update cache for quiz %s: %v", quizID, err)
		return
	}

	quizBytes, err := json.Marshal(quiz)
	if err != nil {
		log.Printf("Failed to marshal quiz %s for cache: %v", quizID, err)
		return
	}

	r.cache.Set(ctx, quizID.String(), string(quizBytes), 0)
}
