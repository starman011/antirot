package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/starman011/antirot/backend/internal/domain"
	"github.com/starman011/antirot/backend/internal/service"
)

type stubPieceRepo struct{}

func (stubPieceRepo) FindByTopics(_ context.Context, _ []string) ([]domain.Piece, error) {
	return []domain.Piece{{
		ID: "p1", Title: "The Vacuum Tube", GapHook: "hook", Topic: "electronics",
		Difficulty: 2, Format: domain.FormatRead, URL: "https://example.org",
		Creator: "A. Human", Source: "example.org",
	}}, nil
}

type stubCheckInRepo struct{}

func (stubCheckInRepo) Save(_ context.Context, c domain.CheckIn) (domain.CheckIn, error) {
	c.ID = "c1"
	c.CreatedAt = time.Now().UTC()
	return c, nil
}

func testRouter() http.Handler {
	return NewRouter(
		service.NewSessionService(stubPieceRepo{}),
		service.NewCheckInService(stubCheckInRepo{}),
	)
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

func TestCheckInCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/checkins", strings.NewReader(`{"state":"restless"}`))
	testRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("status = %d, want 201; body: %s", rec.Code, rec.Body.String())
	}
}

func TestCheckInInvalidState(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/checkins", strings.NewReader(`{"state":"angry"}`))
	testRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 400", rec.Code)
	}
}

func TestInterpret(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/interpret", strings.NewReader(`{"text":"cant stop scrolling reels"}`))
	testRouter().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200; body: %s", rec.Code, rec.Body.String())
	}
	var got struct {
		State string `json:"state"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if got.State != "doomscrolling" {
		t.Errorf("state = %q, want doomscrolling", got.State)
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
