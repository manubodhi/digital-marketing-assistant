package response

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()
	payload := map[string]string{"message": "success"}

	JSON(w, http.StatusOK, payload)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var res APIResponse
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !res.Success {
		t.Errorf("expected success true, got false")
	}
}

func TestError(t *testing.T) {
	w := httptest.NewRecorder()

	Error(w, http.StatusBadRequest, "INVALID_INPUT", "Input field is required")

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}

	var res APIResponse
	if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
		t.Fatalf("failed to decode error response: %v", err)
	}

	if res.Success {
		t.Errorf("expected success false, got true")
	}
	if res.Error == nil || res.Error.Code != "INVALID_INPUT" {
		t.Errorf("expected error code INVALID_INPUT, got %v", res.Error)
	}
}
