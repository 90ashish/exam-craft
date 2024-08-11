package services

import (
	"context"
	"errors"

	"exam-craft/internal/auth"
	"exam-craft/internal/models"
	"exam-craft/internal/pkg/utils"
	"exam-craft/internal/repositories"

	"github.com/google/uuid"
)

// UserService handles business logic for users
type UserService struct {
	userRepo *repositories.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// RegisterUser registers a new user
func (s *UserService) RegisterUser(ctx context.Context, user *models.User) error {
	user.ID = uuid.New()
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.userRepo.CreateUser(ctx, user)
}

// LoginUser logs in a user
func (s *UserService) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

// GetOrCreateUserByGoogleID retrieves a user by their Google ID, or creates one if they don't exist.
func (s *UserService) GetOrCreateUserByGoogleID(userInfo auth.GoogleUserInfo) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(context.Background(), userInfo.Email)
	if err == nil {
		// User exists, return user
		return user, nil
	}

	// If user does not exist, create a new one
	newUser := &models.User{
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Password: "",        // No password needed for Google OAuth users
		Role:     "student", // Default role, could be customized
	}

	err = s.userRepo.CreateUser(context.Background(), newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
