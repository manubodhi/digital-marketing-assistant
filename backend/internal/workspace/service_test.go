package workspace

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database/db"
)

type MockWorkspaceRepo struct {
	workspaces map[uuid.UUID]*db.Workspace
	members    map[string]*db.WorkspaceMember // key: workspaceID:userID
}

func NewMockWorkspaceRepo() *MockWorkspaceRepo {
	return &MockWorkspaceRepo{
		workspaces: make(map[uuid.UUID]*db.Workspace),
		members:    make(map[string]*db.WorkspaceMember),
	}
}

func (m *MockWorkspaceRepo) CreateWorkspace(ctx context.Context, id uuid.UUID, name, slug string) (*db.Workspace, error) {
	ws := &db.Workspace{
		ID:        database.UUIDToPgtype(id),
		Name:      name,
		Slug:      slug,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	m.workspaces[id] = ws
	return ws, nil
}

func (m *MockWorkspaceRepo) GetWorkspaceByID(ctx context.Context, id uuid.UUID) (*db.Workspace, error) {
	ws, ok := m.workspaces[id]
	if !ok {
		return nil, nil
	}
	return ws, nil
}

func (m *MockWorkspaceRepo) GetWorkspaceBySlug(ctx context.Context, slug string) (*db.Workspace, error) {
	for _, ws := range m.workspaces {
		if ws.Slug == slug {
			return ws, nil
		}
	}
	return nil, nil
}

func (m *MockWorkspaceRepo) AddWorkspaceMember(ctx context.Context, workspaceID, userID uuid.UUID, role string) (*db.WorkspaceMember, error) {
	key := workspaceID.String() + ":" + userID.String()
	member := &db.WorkspaceMember{
		WorkspaceID: database.UUIDToPgtype(workspaceID),
		UserID:      database.UUIDToPgtype(userID),
		Role:        role,
		JoinedAt:    pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	m.members[key] = member
	return member, nil
}

func (m *MockWorkspaceRepo) ListUserWorkspaces(ctx context.Context, userID uuid.UUID) ([]db.ListUserWorkspacesRow, error) {
	var rows []db.ListUserWorkspacesRow
	for _, member := range m.members {
		if database.PgtypeToUUID(member.UserID) == userID {
			ws := m.workspaces[database.PgtypeToUUID(member.WorkspaceID)]
			if ws != nil {
				rows = append(rows, db.ListUserWorkspacesRow{
					ID:        ws.ID,
					Name:      ws.Name,
					Slug:      ws.Slug,
					Role:      member.Role,
					CreatedAt: ws.CreatedAt,
				})
			}
		}
	}
	return rows, nil
}

func (m *MockWorkspaceRepo) GetWorkspaceMembership(ctx context.Context, workspaceID, userID uuid.UUID) (*db.WorkspaceMember, error) {
	key := workspaceID.String() + ":" + userID.String()
	member, ok := m.members[key]
	if !ok {
		return nil, nil
	}
	return member, nil
}

func TestWorkspaceService_CreateWorkspace(t *testing.T) {
	repo := NewMockWorkspaceRepo()
	svc := NewService(repo)

	ctx := context.Background()
	userID, _ := uuid.NewV7()

	ws, err := svc.CreateWorkspace(ctx, userID, "Acme Marketing", "")
	if err != nil {
		t.Fatalf("expected successful creation, got %v", err)
	}

	if ws.Name != "Acme Marketing" {
		t.Errorf("expected name Acme Marketing, got %s", ws.Name)
	}
	if ws.Slug != "acme-marketing" {
		t.Errorf("expected slug acme-marketing, got %s", ws.Slug)
	}
	if ws.Role != "owner" {
		t.Errorf("expected role owner, got %s", ws.Role)
	}
}

func TestWorkspaceService_ListUserWorkspaces(t *testing.T) {
	repo := NewMockWorkspaceRepo()
	svc := NewService(repo)

	ctx := context.Background()
	userID, _ := uuid.NewV7()

	_, _ = svc.CreateWorkspace(ctx, userID, "Workspace 1", "")
	_, _ = svc.CreateWorkspace(ctx, userID, "Workspace 2", "")

	list, err := svc.ListUserWorkspaces(ctx, userID)
	if err != nil {
		t.Fatalf("expected successful list, got %v", err)
	}

	if len(list) != 2 {
		t.Errorf("expected 2 workspaces, got %d", len(list))
	}
}

func TestWorkspaceService_VerifyMembership(t *testing.T) {
	repo := NewMockWorkspaceRepo()
	svc := NewService(repo)

	ctx := context.Background()
	userID, _ := uuid.NewV7()
	otherUserID, _ := uuid.NewV7()

	ws, _ := svc.CreateWorkspace(ctx, userID, "Private Workspace", "")

	// Valid member
	role, err := svc.VerifyMembership(ctx, ws.ID, userID)
	if err != nil || role != "owner" {
		t.Errorf("expected owner role, got %s, err: %v", role, err)
	}

	// Unauthorized user
	_, err = svc.VerifyMembership(ctx, ws.ID, otherUserID)
	if err != ErrUnauthorized {
		t.Errorf("expected ErrUnauthorized, got %v", err)
	}
}
