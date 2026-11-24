package workspace

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

func (r *SQLRepository) CreateWorkspace(ctx context.Context, id uuid.UUID, name, slug string) (*db.Workspace, error) {
	ws, err := r.queries.CreateWorkspace(ctx, db.CreateWorkspaceParams{
		ID:   database.UUIDToPgtype(id),
		Name: name,
		Slug: slug,
	})
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

func (r *SQLRepository) GetWorkspaceByID(ctx context.Context, id uuid.UUID) (*db.Workspace, error) {
	ws, err := r.queries.GetWorkspaceByID(ctx, database.UUIDToPgtype(id))
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

func (r *SQLRepository) GetWorkspaceBySlug(ctx context.Context, slug string) (*db.Workspace, error) {
	ws, err := r.queries.GetWorkspaceBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

func (r *SQLRepository) AddWorkspaceMember(ctx context.Context, workspaceID, userID uuid.UUID, role string) (*db.WorkspaceMember, error) {
	m, err := r.queries.AddWorkspaceMember(ctx, db.AddWorkspaceMemberParams{
		WorkspaceID: database.UUIDToPgtype(workspaceID),
		UserID:      database.UUIDToPgtype(userID),
		Role:        role,
	})
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *SQLRepository) ListUserWorkspaces(ctx context.Context, userID uuid.UUID) ([]db.ListUserWorkspacesRow, error) {
	return r.queries.ListUserWorkspaces(ctx, database.UUIDToPgtype(userID))
}

func (r *SQLRepository) GetWorkspaceMembership(ctx context.Context, workspaceID, userID uuid.UUID) (*db.WorkspaceMember, error) {
	m, err := r.queries.GetWorkspaceMembership(ctx, db.GetWorkspaceMembershipParams{
		WorkspaceID: database.UUIDToPgtype(workspaceID),
		UserID:      database.UUIDToPgtype(userID),
	})
	if err != nil {
		return nil, err
	}
	return &m, nil
}
