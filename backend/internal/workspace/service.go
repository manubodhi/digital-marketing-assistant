package workspace

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database/db"
)

var (
	ErrWorkspaceNotFound = errors.New("workspace not found")
	ErrInvalidSlug       = errors.New("invalid workspace slug")
	ErrUnauthorized      = errors.New("unauthorized to access this workspace")
)

type WorkspaceDTO struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Role      string    `json:"role"`
	CreatedAt string    `json:"created_at"`
}

type Repository interface {
	CreateWorkspace(ctx context.Context, id uuid.UUID, name, slug string) (*db.Workspace, error)
	GetWorkspaceByID(ctx context.Context, id uuid.UUID) (*db.Workspace, error)
	GetWorkspaceBySlug(ctx context.Context, slug string) (*db.Workspace, error)
	AddWorkspaceMember(ctx context.Context, workspaceID, userID uuid.UUID, role string) (*db.WorkspaceMember, error)
	ListUserWorkspaces(ctx context.Context, userID uuid.UUID) ([]db.ListUserWorkspacesRow, error)
	GetWorkspaceMembership(ctx context.Context, workspaceID, userID uuid.UUID) (*db.WorkspaceMember, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

var nonAlphaNum = regexp.MustCompile(`[^a-z0-9]+`)

func Slugify(name string) string {
	slug := strings.ToLower(strings.TrimSpace(name))
	slug = nonAlphaNum.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	if slug == "" {
		slug = "workspace"
	}
	return slug
}

func (s *Service) CreateWorkspace(ctx context.Context, userID uuid.UUID, name, customSlug string) (*WorkspaceDTO, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("workspace name is required")
	}

	slug := customSlug
	if slug == "" {
		slug = Slugify(name)
	} else {
		slug = Slugify(slug)
	}

	// Check slug uniqueness; append random suffix if collision occurs
	existing, err := s.repo.GetWorkspaceBySlug(ctx, slug)
	if err == nil && existing != nil {
		slug = fmt.Sprintf("%s-%s", slug, uuid.New().String()[:6])
	}

	workspaceID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to generate workspace uuid: %w", err)
	}

	ws, err := s.repo.CreateWorkspace(ctx, workspaceID, name, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to create workspace: %w", err)
	}

	// Creator automatically becomes 'owner'
	_, err = s.repo.AddWorkspaceMember(ctx, workspaceID, userID, "owner")
	if err != nil {
		return nil, fmt.Errorf("failed to add workspace owner: %w", err)
	}

	return &WorkspaceDTO{
		ID:        workspaceID,
		Name:      ws.Name,
		Slug:      ws.Slug,
		Role:      "owner",
		CreatedAt: ws.CreatedAt.Time.String(),
	}, nil
}

func (s *Service) ListUserWorkspaces(ctx context.Context, userID uuid.UUID) ([]WorkspaceDTO, error) {
	rows, err := s.repo.ListUserWorkspaces(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to list workspaces: %w", err)
	}

	result := make([]WorkspaceDTO, 0, len(rows))
	for _, r := range rows {
		result = append(result, WorkspaceDTO{
			ID:        r.ID.Bytes,
			Name:      r.Name,
			Slug:      r.Slug,
			Role:      r.Role,
			CreatedAt: r.CreatedAt.Time.String(),
		})
	}

	return result, nil
}

func (s *Service) VerifyMembership(ctx context.Context, workspaceID, userID uuid.UUID) (string, error) {
	member, err := s.repo.GetWorkspaceMembership(ctx, workspaceID, userID)
	if err != nil || member == nil {
		return "", ErrUnauthorized
	}
	return member.Role, nil
}
