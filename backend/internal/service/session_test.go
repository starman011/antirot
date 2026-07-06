package service

import (
	"context"
	"errors"
	"testing"

	"github.com/starman011/antirot/backend/internal/domain"
)

type stubPieceRepo struct {
	pieces []domain.Piece
	err    error
}

func (s *stubPieceRepo) FindByTopics(_ context.Context, _ []string) ([]domain.Piece, error) {
	return s.pieces, s.err
}

func TestPickPieceReturnsOne(t *testing.T) {
	repo := &stubPieceRepo{pieces: []domain.Piece{
		{ID: "p1", Title: "The Vacuum Tube", Creator: "A. Human", Source: "example.org"},
		{ID: "p2", Title: "Second Piece", Creator: "B. Human", Source: "example.org"},
	}}
	svc := NewSessionService(repo)

	got, err := svc.PickPiece(context.Background(), domain.StateDoomscrolling, []string{"electronics"})
	if err != nil {
		t.Fatalf("PickPiece() error = %v", err)
	}
	if got.ID != "p1" && got.ID != "p2" {
		t.Errorf("PickPiece() ID = %q, want one of the candidates", got.ID)
	}
}

func TestPickPieceInvalidState(t *testing.T) {
	svc := NewSessionService(&stubPieceRepo{})
	if _, err := svc.PickPiece(context.Background(), domain.State("angry"), nil); err == nil {
		t.Error("PickPiece() with invalid state: want error, got nil")
	}
}

func TestPickPieceEmptyCuration(t *testing.T) {
	svc := NewSessionService(&stubPieceRepo{})
	_, err := svc.PickPiece(context.Background(), domain.StateRestless, nil)
	if !errors.Is(err, ErrNoPiece) {
		t.Errorf("PickPiece() error = %v, want ErrNoPiece", err)
	}
}
