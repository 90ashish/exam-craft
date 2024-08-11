package repositories

import (
	"context"
	"exam-craft/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// StudentResponseRepository handles database operations for student responses
type StudentResponseRepository struct {
	db *sqlx.DB
}

// NewStudentResponseRepository creates a new StudentResponseRepository
func NewStudentResponseRepository(db *sqlx.DB) *StudentResponseRepository {
	return &StudentResponseRepository{db: db}
}

// CreateStudentResponse inserts a new student response into the database
func (r *StudentResponseRepository) CreateStudentResponse(ctx context.Context, response *models.StudentResponse) error {
	response.ID = uuid.New()
	query := `INSERT INTO student_responses (id, student_id, quiz_id, responses, submitted_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, response.ID, response.StudentID, response.QuizID, pq.Array(response.Responses), response.SubmittedAt)
	return err
}

// GetResponsesByQuizID retrieves student responses by quiz ID from the database
func (r *StudentResponseRepository) GetResponsesByQuizID(ctx context.Context, quizID uuid.UUID) ([]models.StudentResponse, error) {
	var responses []models.StudentResponse
	query := `SELECT * FROM student_responses WHERE quiz_id = $1`
	err := r.db.SelectContext(ctx, &responses, query, quizID)
	return responses, err
}
