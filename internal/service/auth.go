package service

import (
	"context"
	"errors"

	"byte-cabinet/internal/model"
	"byte-cabinet/internal/repository"
	"byte-cabinet/pkg/utils"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidOldPassword = errors.New("invalid old password")
)

// AuthService handles authentication business logic
type AuthService struct {
	userRepo   *repository.UserRepository
	jwtManager *utils.JWTManager
}

// NewAuthService creates a new authentication service
func NewAuthService(userRepo *repository.UserRepository, jwtManager *utils.JWTManager) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

// LoginResult represents the result of a successful login
type LoginResult struct {
	User   *model.UserResponse `json:"user"`
	Tokens *utils.TokenPair    `json:"tokens"`
}

// Login authenticates a user and returns tokens
func (s *AuthService) Login(ctx context.Context, req *model.UserLogin) (*LoginResult, error) {
	// Find user by username
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify password
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// Generate tokens
	tokens, err := s.jwtManager.GenerateTokenPair(user.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		User:   user.ToResponse(),
		Tokens: tokens,
	}, nil
}

// RefreshTokens generates new tokens using a refresh token
func (s *AuthService) RefreshTokens(ctx context.Context, refreshToken string) (*utils.TokenPair, error) {
	// Validate refresh token
	claims, err := s.jwtManager.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// Verify user still exists
	user, err := s.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Generate new tokens
	return s.jwtManager.GenerateTokenPair(user.ID, user.Username)
}

// GetCurrentUser retrieves the current user by ID
func (s *AuthService) GetCurrentUser(ctx context.Context, userID int64) (*model.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user.ToResponse(), nil
}

// ChangePassword changes the user's password
func (s *AuthService) ChangePassword(ctx context.Context, userID int64, req *model.PasswordChange) error {
	// Get current user
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return ErrUserNotFound
		}
		return err
	}

	// Verify old password
	if !utils.CheckPassword(req.OldPassword, user.PasswordHash) {
		return ErrInvalidOldPassword
	}

	// Hash new password
	newHash, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	// Update password
	return s.userRepo.UpdatePassword(ctx, userID, newHash)
}

// UpdateProfile updates the user's profile
func (s *AuthService) UpdateProfile(ctx context.Context, userID int64, req *model.UserUpdate) (*model.UserResponse, error) {
	// Verify user exists
	_, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Update profile
	if err := s.userRepo.UpdateProfile(ctx, userID, req); err != nil {
		return nil, err
	}

	// Get updated user
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}
