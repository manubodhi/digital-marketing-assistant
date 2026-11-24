package config

import (
	"os"
	"testing"
)

func TestLoad_Defaults(t *testing.T) {
	// Clear any existing env vars for this test
	os.Unsetenv("PORT")
	os.Unsetenv("DATABASE_URL")
	os.Unsetenv("JWT_SECRET")
	os.Unsetenv("APP_ENV")

	cfg := Load()

	if cfg.Port != "8080" {
		t.Errorf("expected default port 8080, got %s", cfg.Port)
	}
	if cfg.Environment != "development" {
		t.Errorf("expected default env development, got %s", cfg.Environment)
	}
	if cfg.DatabaseURL == "" {
		t.Errorf("expected non-empty default DatabaseURL")
	}
	if cfg.JWTSecret == "" {
		t.Errorf("expected non-empty default JWTSecret")
	}
}

func TestLoad_CustomEnv(t *testing.T) {
	os.Setenv("PORT", "9090")
	os.Setenv("APP_ENV", "production")
	defer func() {
		os.Unsetenv("PORT")
		os.Unsetenv("APP_ENV")
	}()

	cfg := Load()

	if cfg.Port != "9090" {
		t.Errorf("expected port 9090, got %s", cfg.Port)
	}
	if cfg.Environment != "production" {
		t.Errorf("expected env production, got %s", cfg.Environment)
	}
}
