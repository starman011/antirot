// Package repository owns all data access and returns domain models, never raw rows.
package repository

import (
	"context"
	"strings"
	"sync"

	"github.com/starman011/antirot/backend/internal/domain"
)

type MemoryPieceRepository struct {
	mu     sync.RWMutex
	pieces []domain.Piece
}

// NewMemoryPieceRepository is the Phase-0/1 stand-in; Postgres replaces it in T104.
func NewMemoryPieceRepository() *MemoryPieceRepository {
	return &MemoryPieceRepository{
		pieces: []domain.Piece{
			{
				ID:         "piece-vacuum-tube",
				Title:      "The Vacuum Tube",
				GapHook:    "Ever wondered why vacuum tubes are back in high-end audio?",
				Topic:      "electronics",
				Difficulty: 2,
				Format:     domain.FormatRead,
				URL:        "https://example.org/vacuum-tube",
				Creator:    "Placeholder Human",
				Source:     "example.org",
			},
		},
	}
}

func (r *MemoryPieceRepository) FindByTopics(_ context.Context, topics []string) ([]domain.Piece, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(topics) == 0 {
		out := make([]domain.Piece, len(r.pieces))
		copy(out, r.pieces)
		return out, nil
	}

	want := make(map[string]bool, len(topics))
	for _, t := range topics {
		want[strings.ToLower(strings.TrimSpace(t))] = true
	}

	var out []domain.Piece
	for _, p := range r.pieces {
		if want[strings.ToLower(p.Topic)] {
			out = append(out, p)
		}
	}
	return out, nil
}
