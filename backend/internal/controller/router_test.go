package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/starman011/antirot/backend/internal/repository"
	"github.com/starman011/antirot/backend/internal/service"
)

func testRouter() http.Handler {
	return NewRouter(service.NewSessionService(repository.NewMemoryPieceRepository()))
}

func TestHealth(t *testing.T) {
	rec := httptest.NewRecorder()
	testRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/health", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}

func TestSessionPiece(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/session/piece?state=doomscrolling&interests=electronics", nil)
	testRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body: %s", rec.Code, rec.Body.String())
	}

	var got struct {
		Creator string `json:"creator"`
		Source  string `json:"source"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	// Principle II: provenance must always be surfaced.
	if got.Creator == "" || got.Source == "" {
		t.Errorf("piece missing provenance: creator=%q source=%q", got.Creator, got.Source)
	}
}

func TestSessionPieceInvalidState(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/session/piece?state=angry", nil)
	testRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestSessionPieceSkippedCheckIn(t *testing.T) {
	// Check-in is skippable (Principle IV): no state param must still work.
	rec := httptest.NewRecorder()
	testRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/session/piece", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body: %s", rec.Code, rec.Body.String())
	}
}
