package handlers

import (
	"context"
	"exam-craft/internal/models"
	"exam-craft/internal/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// StudentHandler handles student-related requests
type StudentHandler struct {
	studentService *services.StudentService
}

// NewStudentHandler creates a new StudentHandler
func NewStudentHandler(studentService *services.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

// ListQuizzes handles listing available quizzes for students
func (h *StudentHandler) ListQuizzes(c *gin.Context) {
	quizzes, err := h.studentService.ListAvailableQuizzes(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quizzes)
}

// SubmitResponse handles submitting a student's response to a quiz
func (h *StudentHandler) SubmitResponse(c *gin.Context) {
	var response models.StudentResponse
	if err := c.BindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentID := c.MustGet("userID").(uuid.UUID) // Extract from middleware
	response.StudentID = studentID
	response.SubmittedAt = time.Now().Unix()

	if err := h.studentService.SubmitResponse(context.Background(), &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}
