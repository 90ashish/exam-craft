package handlers

import (
	"context"
	"exam-craft/internal/models"
	"exam-craft/internal/pkg/utils"
	"exam-craft/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// QuizHandler handles quiz-related requests
type QuizHandler struct {
	quizService     *services.QuizService
	questionService *services.QuestionService
}

// NewQuizHandler creates a new QuizHandler
func NewQuizHandler(quizService *services.QuizService, questionService *services.QuestionService) *QuizHandler {
	return &QuizHandler{quizService: quizService, questionService: questionService}
}

// CreateQuiz handles creating a new quiz
func (h *QuizHandler) CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.BindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateQuizTitle(quiz.Title); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	instructorID := c.MustGet("userID").(uuid.UUID) // Extract from middleware
	quiz.InstructorID = instructorID

	if err := h.quizService.CreateQuiz(context.Background(), &quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, quiz)
}

// GetQuiz handles retrieving a quiz by ID
func (h *QuizHandler) GetQuiz(c *gin.Context) {
	quizID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quiz ID"})
		return
	}

	quiz, err := h.quizService.GetQuiz(context.Background(), quizID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "quiz not found"})
		return
	}

	c.JSON(http.StatusOK, quiz)
}

// AddQuestion handles adding a question to a quiz
func (h *QuizHandler) AddQuestion(c *gin.Context) {
	var question models.Question
	if err := c.BindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateQuestionType(question.Type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.questionService.CreateQuestion(context.Background(), &question); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, question)
}
