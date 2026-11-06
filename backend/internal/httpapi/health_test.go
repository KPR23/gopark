package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth_OK(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected content type %s, got %s", "application/json", rec.Header().Get("Content-Type"))
	}
	var gotHealthResponse healthResponse
	if err := json.NewDecoder(rec.Body).Decode(&gotHealthResponse); err != nil {
		t.Fatalf("decode response body: %v", err)
	}
	if gotHealthResponse.Status != "ok" {
		t.Errorf("expected status %s, got %s", "ok", gotHealthResponse.Status)
	}
}

func TestHealth_MethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	NewRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, rec.Code)
	}
}
