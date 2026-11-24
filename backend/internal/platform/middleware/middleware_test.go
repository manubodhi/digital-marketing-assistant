package middleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/token"
)

type mockVerifier struct {
	allowed map[string]string // key: workspaceID:userID -> role
}

func (m *mockVerifier) VerifyMembership(ctx context.Context, workspaceID, userID uuid.UUID) (string, error) {
	key := workspaceID.String() + ":" + userID.String()
	role, ok := m.allowed[key]
	if !ok {
		return "", errors.New("unauthorized")
	}
	return role, nil
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	tm := token.NewTokenManager("test-secret-key-32-bytes-long!!", 15*time.Minute)
	userID, _ := uuid.NewV7()
	tok, _ := tm.GenerateAccessToken(userID, "test@example.com")

	handler := Auth(tm)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := GetUserID(r.Context())
		if !ok || id != userID {
			t.Errorf("expected user id %v, got %v", userID, id)
		}
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tok)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}

func TestAuthMiddleware_MissingToken(t *testing.T) {
	tm := token.NewTokenManager("test-secret-key-32-bytes-long!!", 15*time.Minute)

	handler := Auth(tm)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest("GET", "/protected", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", rec.Code)
	}
}

func TestTenantGuardMiddleware(t *testing.T) {
	userID, _ := uuid.NewV7()
	validWS, _ := uuid.NewV7()
	invalidWS, _ := uuid.NewV7()

	verifier := &mockVerifier{
		allowed: map[string]string{
			validWS.String() + ":" + userID.String(): "owner",
		},
	}

	handler := TenantGuard(verifier)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wsID, ok := GetTenantID(r.Context())
		role, _ := GetTenantRole(r.Context())
		if !ok || wsID != validWS || role != "owner" {
			t.Errorf("context tenant mismatch")
		}
		w.WriteHeader(http.StatusOK)
	}))

	// Case 1: Valid workspace access
	req := httptest.NewRequest("GET", "/workspace-data", nil)
	ctx := context.WithValue(req.Context(), UserIDKey, userID)
	req = req.WithContext(ctx)
	req.Header.Set("X-Workspace-ID", validWS.String())
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	// Case 2: Forbidden workspace access (cross-tenant leak prevention)
	reqForbidden := httptest.NewRequest("GET", "/workspace-data", nil)
	ctxForbidden := context.WithValue(reqForbidden.Context(), UserIDKey, userID)
	reqForbidden = reqForbidden.WithContext(ctxForbidden)
	reqForbidden.Header.Set("X-Workspace-ID", invalidWS.String())
	recForbidden := httptest.NewRecorder()

	handler.ServeHTTP(recForbidden, reqForbidden)
	if recForbidden.Code != http.StatusForbidden {
		t.Errorf("expected status 403, got %d", recForbidden.Code)
	}
}
