package auth

import (
	"encoding/json"
	"net/http"

	"github.com/manubodhi/digital-marketing-assistant/backend/internal/platform/response"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_REQUEST", "Failed to parse request body")
		return
	}

	result, err := h.service.Register(r.Context(), req.Email, req.Password, req.FullName)
	if err != nil {
		switch err {
		case ErrUserAlreadyExists:
			response.Error(w, http.StatusConflict, "USER_EXISTS", err.Error())
		case ErrPasswordTooShort, ErrInvalidEmail:
			response.Error(w, http.StatusBadRequest, "VALIDATION_FAILED", err.Error())
		default:
			response.Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Registration failed")
		}
		return
	}

	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_REQUEST", "Failed to parse request body")
		return
	}

	result, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if err == ErrInvalidCredentials {
			response.Error(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid email or password")
			return
		}
		response.Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Login failed")
		return
	}

	response.JSON(w, http.StatusOK, result)
}
