// Package service holds business logic; repository interfaces are defined here (consumer side).
package service

import (
	"context"
	"fmt"
	"math/rand/v2"

	"github.com/starman011/antirot/backend/internal/domain"
)

type PieceRepository interface {
	// Empty topics means no filter.
	FindByTopics(ctx context.Context, topics []string) ([]domain.Piece, error)
}

// SessionService returns one piece per session; the loop is finite by design (Principle I).
type SessionService struct {
	pieces PieceRepository
}

func NewSessionService(pieces PieceRepository) *SessionService {
	return &SessionService{pieces: pieces}
}

// PickPiece picks uniformly among matches for the MVP; taste-graph ranking arrives in Phase 3.
func (s *SessionService) PickPiece(ctx context.Context, state domain.State, interests []string) (domain.Piece, error) {
	if !state.Valid() {
		return domain.Piece{}, fmt.Errorf("session: invalid state %q", state)
	}

	candidates, err := s.pieces.FindByTopics(ctx, interests)
	if err != nil {
		return domain.Piece{}, fmt.Errorf("session: find pieces: %w", err)
	}
	if len(candidates) == 0 {
		return domain.Piece{}, ErrNoPiece
	}
	return candidates[rand.IntN(len(candidates))], nil
}

var ErrNoPiece = fmt.Errorf("session: no curated piece available")
