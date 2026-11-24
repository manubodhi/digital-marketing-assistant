package auth

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/crypto"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database/db"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/token"
)

var (
	ErrUserAlreadyExists = errors.New("user with this email already exists")
	ErrInvalidCredentials = errors.New("invalid email or password")
	ErrPasswordTooShort   = errors.New("password must be at least 8 characters long")
	ErrInvalidEmail       = errors.New("invalid email address")
)

type UserDTO struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	FullName        string    `json:"full_name"`
	IsEmailVerified bool      `json:"is_email_verified"`
	TotpEnabled     bool      `json:"totp_enabled"`
}

type AuthResult struct {
	AccessToken string  `json:"access_token"`
	User        UserDTO `json:"user"`
}

type Repository interface {
	CreateUser(ctx context.Context, id uuid.UUID, email, passwordHash, fullName string) (*db.User, error)
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*db.User, error)
}

type Service struct {
	repo         Repository
	tokenManager *token.TokenManager
}

func NewService(repo Repository, tokenManager *token.TokenManager) *Service {
	return &Service{
		repo:         repo,
		tokenManager: tokenManager,
	}
}

func (s *Service) Register(ctx context.Context, email, password, fullName string) (*AuthResult, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	fullName = strings.TrimSpace(fullName)

	if email == "" || !strings.Contains(email, "@") {
		return nil, ErrInvalidEmail
	}
	if len(password) < 8 {
		return nil, ErrPasswordTooShort
	}

	// Check if user already exists
	existing, err := s.repo.GetUserByEmail(ctx, email)
	if err == nil && existing != nil {
		return nil, ErrUserAlreadyExists
	}

	// Hash password using Argon2id
	hashedPassword, err := crypto.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Generate UUIDv7 for sequential index locality
	userID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuidv7: %w", err)
	}

	user, err := s.repo.CreateUser(ctx, userID, email, hashedPassword, fullName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate Access Token (15m)
	accessToken, err := s.tokenManager.GenerateAccessToken(userID, email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &AuthResult{
		AccessToken: accessToken,
		User: UserDTO{
			ID:              userID,
			Email:           user.Email,
			FullName:        user.FullName,
			IsEmailVerified: user.IsEmailVerified,
			TotpEnabled:     user.TotpEnabled,
		},
	}, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (*AuthResult, error) {
	email = strings.TrimSpace(strings.ToLower(email))

	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, ErrInvalidCredentials
	}

	if !user.PasswordHash.Valid {
		return nil, ErrInvalidCredentials
	}

	// Constant-time Argon2id comparison
	valid, err := crypto.ComparePassword(password, user.PasswordHash.String)
	if err != nil || !valid {
		return nil, ErrInvalidCredentials
	}

	userID := user.ID.Bytes

	accessToken, err := s.tokenManager.GenerateAccessToken(userID, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	return &AuthResult{
		AccessToken: accessToken,
		User: UserDTO{
			ID:              userID,
			Email:           user.Email,
			FullName:        user.FullName,
			IsEmailVerified: user.IsEmailVerified,
			TotpEnabled:     user.TotpEnabled,
		},
	}, nil
}
