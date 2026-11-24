package workspace

import (
	"encoding/json"
	"net/http"

	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/middleware"
	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/response"
)

type CreateWorkspaceRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "User not authenticated")
		return
	}

	var req CreateWorkspaceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_REQUEST", "Failed to parse request body")
		return
	}

	ws, err := h.service.CreateWorkspace(r.Context(), userID, req.Name, req.Slug)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "CREATION_FAILED", err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, ws)
}

func (h *Handler) ListWorkspaces(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r.Context())
	if !ok {
		response.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "User not authenticated")
		return
	}

	list, err := h.service.ListUserWorkspaces(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Failed to retrieve workspaces")
		return
	}

	response.JSON(w, http.StatusOK, list)
}
