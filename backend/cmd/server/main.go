package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/auth"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/config"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/database/db"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/middleware"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/response"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/token"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/workspace"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	// 1. Database Connection Pool
	pool, err := database.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Fatal: Failed to initialize database: %v", err)
	}
	defer pool.Close()

	// 2. sqlc Queries Engine
	queries := db.New(pool)

	// 3. Platform Managers
	tokenManager := token.NewTokenManager(cfg.JWTSecret, 15*time.Minute)

	// 4. Domain Repositories & Services (Constructor Factory Pattern)
	authRepo := auth.NewRepository(queries)
	authService := auth.NewService(authRepo, tokenManager)
	authHandler := auth.NewHandler(authService)

	workspaceRepo := workspace.NewRepository(queries)
	workspaceService := workspace.NewService(workspaceRepo)
	workspaceHandler := workspace.NewHandler(workspaceService)

	// 5. HTTP Router
	r := chi.NewRouter()

	r.Use(chimw.RequestID)
	r.Use(chimw.RealIP)
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.Timeout(60 * time.Second))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-Workspace-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health Check
	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now().UTC().Format(time.RFC3339),
			"version":   "0.1.0",
		})
	})

	// API v1 Routes
	r.Route("/api/v1", func(r chi.Router) {
		// Public Auth Routes
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", authHandler.Register)
			r.Post("/login", authHandler.Login)
		})

		// Protected Workspace Routes
		r.Group(func(r chi.Router) {
			r.Use(middleware.Auth(tokenManager))

			r.Route("/workspaces", func(r chi.Router) {
				r.Get("/", workspaceHandler.ListWorkspaces)
				r.Post("/", workspaceHandler.CreateWorkspace)
			})
		})
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server listening on port %s (env: %s)...", cfg.Port, cfg.Environment)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
