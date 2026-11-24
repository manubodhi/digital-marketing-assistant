package auth

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/crypto"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database/db"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/token"
)

type MockAuthRepo struct {
	users map[string]*db.User
}

func NewMockAuthRepo() *MockAuthRepo {
	return &MockAuthRepo{
		users: make(map[string]*db.User),
	}
}

func (m *MockAuthRepo) CreateUser(ctx context.Context, id uuid.UUID, email, passwordHash, fullName string) (*db.User, error) {
	u := &db.User{
		ID:              database.UUIDToPgtype(id),
		Email:           email,
		PasswordHash:    database.TextToPgtype(passwordHash),
		FullName:        fullName,
		IsEmailVerified: false,
		TotpEnabled:     false,
	}
	m.users[email] = u
	return u, nil
}

func (m *MockAuthRepo) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	u, ok := m.users[email]
	if !ok {
		return nil, nil
	}
	return u, nil
}

func (m *MockAuthRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*db.User, error) {
	for _, u := range m.users {
		if database.PgtypeToUUID(u.ID) == id {
			return u, nil
		}
	}
	return nil, nil
}

func TestAuthService_Register_Success(t *testing.T) {
	repo := NewMockAuthRepo()
	tm := token.NewTokenManager("test-secret-key-min-32-bytes-long!", 15*time.Minute)
	svc := NewService(repo, tm)

	ctx := context.Background()
	res, err := svc.Register(ctx, "test@example.com", "securePassword123", "Test User")
	if err != nil {
		t.Fatalf("expected successful registration, got: %v", err)
	}

	if res.AccessToken == "" {
		t.Errorf("expected access token to be generated")
	}
	if res.User.Email != "test@example.com" {
		t.Errorf("expected email test@example.com, got %s", res.User.Email)
	}
}

func TestAuthService_Register_DuplicateEmail(t *testing.T) {
	repo := NewMockAuthRepo()
	tm := token.NewTokenManager("test-secret-key-min-32-bytes-long!", 15*time.Minute)
	svc := NewService(repo, tm)

	ctx := context.Background()
	_, _ = svc.Register(ctx, "test@example.com", "securePassword123", "Test User")

	_, err := svc.Register(ctx, "test@example.com", "anotherPassword123", "Duplicate User")
	if err != ErrUserAlreadyExists {
		t.Fatalf("expected ErrUserAlreadyExists, got %v", err)
	}
}

func TestAuthService_Register_InvalidInputs(t *testing.T) {
	repo := NewMockAuthRepo()
	tm := token.NewTokenManager("test-secret-key-min-32-bytes-long!", 15*time.Minute)
	svc := NewService(repo, tm)

	ctx := context.Background()

	// Short password
	_, err := svc.Register(ctx, "test@example.com", "short", "Test User")
	if err != ErrPasswordTooShort {
		t.Errorf("expected ErrPasswordTooShort, got %v", err)
	}

	// Invalid email
	_, err = svc.Register(ctx, "invalid-email", "securePassword123", "Test User")
	if err != ErrInvalidEmail {
		t.Errorf("expected ErrInvalidEmail, got %v", err)
	}
}

func TestAuthService_Login_Success(t *testing.T) {
	repo := NewMockAuthRepo()
	tm := token.NewTokenManager("test-secret-key-min-32-bytes-long!", 15*time.Minute)
	svc := NewService(repo, tm)

	ctx := context.Background()
	password := "correctPassword123"
	hash, _ := crypto.HashPassword(password)
	id, _ := uuid.NewV7()
	_, _ = repo.CreateUser(ctx, id, "login@example.com", hash, "Login User")

	res, err := svc.Login(ctx, "login@example.com", password)
	if err != nil {
		t.Fatalf("expected successful login, got %v", err)
	}

	if res.AccessToken == "" {
		t.Errorf("expected access token to be generated")
	}
	if res.User.Email != "login@example.com" {
		t.Errorf("expected email login@example.com, got %s", res.User.Email)
	}
}

func TestAuthService_Login_InvalidPassword(t *testing.T) {
	repo := NewMockAuthRepo()
	tm := token.NewTokenManager("test-secret-key-min-32-bytes-long!", 15*time.Minute)
	svc := NewService(repo, tm)

	ctx := context.Background()
	hash, _ := crypto.HashPassword("correctPassword123")
	id, _ := uuid.NewV7()
	_, _ = repo.CreateUser(ctx, id, "login@example.com", hash, "Login User")

	_, err := svc.Login(ctx, "login@example.com", "wrongPassword")
	if err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}
