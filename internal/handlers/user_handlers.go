package handlers

import (
	"context"
	"exam-craft/internal/auth"
	"exam-craft/internal/models"
	"exam-craft/internal/pkg/utils"
	"exam-craft/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register handles user registration
func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateEmail(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidatePassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateUserRole(user.Role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user.Password = hashedPassword

	if err := h.userService.RegisterUser(context.Background(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// Login handles user login
func (h *UserHandler) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.ValidateEmail(credentials.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.LoginUser(context.Background(), credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GoogleLogin handles Google OAuth login
func (h *UserHandler) GoogleLogin(c *gin.Context) {
	auth.GoogleLogin(c)
}

// GoogleCallback handles the callback from Google OAuth
func (h *UserHandler) GoogleCallback(c *gin.Context) {
	userInfo, err := auth.GoogleCallback(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetOrCreateUserByGoogleID(*userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create or retrieve user"})
		return
	}

	token, err := auth.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
