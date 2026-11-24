package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/response"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/token"
)

type contextKey string

const (
	UserIDKey    contextKey = "user_id"
	UserEmailKey contextKey = "user_email"
	TenantIDKey  contextKey = "workspace_id"
	TenantRoleKey contextKey = "workspace_role"
)

func Auth(tm *token.TokenManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.Error(w, http.StatusUnauthorized, "MISSING_TOKEN", "Authorization header is required")
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				response.Error(w, http.StatusUnauthorized, "INVALID_TOKEN_FORMAT", "Authorization header must be 'Bearer <token>'")
				return
			}

			claims, err := tm.VerifyToken(parts[1])
			if err != nil {
				response.Error(w, http.StatusUnauthorized, "INVALID_TOKEN", "Token is invalid or expired")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, UserEmailKey, claims.Email)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	id, ok := ctx.Value(UserIDKey).(uuid.UUID)
	return id, ok
}

func GetUserEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(UserEmailKey).(string)
	return email, ok
}
