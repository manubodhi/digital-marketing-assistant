package middleware

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/response"
)

type MembershipVerifier interface {
	VerifyMembership(ctx context.Context, workspaceID, userID uuid.UUID) (string, error)
}

func TenantGuard(verifier MembershipVerifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := GetUserID(r.Context())
			if !ok {
				response.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "Authentication required")
				return
			}

			// Extract workspace ID from URL parameter or X-Workspace-ID header
			wsIDStr := chi.URLParam(r, "workspaceId")
			if wsIDStr == "" {
				wsIDStr = r.Header.Get("X-Workspace-ID")
			}

			if wsIDStr == "" {
				response.Error(w, http.StatusBadRequest, "MISSING_TENANT", "Workspace context is required")
				return
			}

			workspaceID, err := uuid.Parse(wsIDStr)
			if err != nil {
				response.Error(w, http.StatusBadRequest, "INVALID_TENANT_ID", "Invalid workspace UUID")
				return
			}

			role, err := verifier.VerifyMembership(r.Context(), workspaceID, userID)
			if err != nil {
				response.Error(w, http.StatusForbidden, "FORBIDDEN_TENANT_ACCESS", "You do not have access to this workspace")
				return
			}

			ctx := context.WithValue(r.Context(), TenantIDKey, workspaceID)
			ctx = context.WithValue(ctx, TenantRoleKey, role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetTenantID(ctx context.Context) (uuid.UUID, bool) {
	id, ok := ctx.Value(TenantIDKey).(uuid.UUID)
	return id, ok
}

func GetTenantRole(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(TenantRoleKey).(string)
	return role, ok
}
