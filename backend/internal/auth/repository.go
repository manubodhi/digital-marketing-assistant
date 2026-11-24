package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database/db"
)

type SQLRepository struct {
	queries *db.Queries
}

func NewRepository(queries *db.Queries) *SQLRepository {
	return &SQLRepository{queries: queries}
}

func (r *SQLRepository) CreateUser(ctx context.Context, id uuid.UUID, email, passwordHash, fullName string) (*db.User, error) {
	params := db.CreateUserParams{
		ID:              database.UUIDToPgtype(id),
		Email:           email,
		PasswordHash:    database.TextToPgtype(passwordHash),
		FullName:        fullName,
		IsEmailVerified: false,
	}

	u, err := r.queries.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *SQLRepository) GetUserByEmail(ctx context.Context, email string) (*db.User, error) {
	u, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *SQLRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*db.User, error) {
	u, err := r.queries.GetUserByID(ctx, database.UUIDToPgtype(id))
	if err != nil {
		return nil, err
	}
	return &u, nil
}
