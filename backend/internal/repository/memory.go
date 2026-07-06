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
// Seeds are hand-picked human-made pieces; re-verify links before launch (T103).
func NewMemoryPieceRepository() *MemoryPieceRepository {
	return &MemoryPieceRepository{
		pieces: []domain.Piece{
			{
				ID:         "piece-mechanical-watch",
				Title:      "Mechanical Watch",
				GapHook:    "How can a machine with no battery keep time for days, powered by a wrist flick?",
				Topic:      "electronics",
				Difficulty: 3,
				Format:     domain.FormatRead,
				URL:        "https://ciechanow.ski/mechanical-watch/",
				Creator:    "Bartosz Ciechanowski",
				Source:     "ciechanow.ski",
			},
			{
				ID:         "piece-vacuum-tubes",
				Title:      "How Vacuum Tubes Work",
				GapHook:    "Ever wondered why vacuum tubes are back in high-end audio?",
				Topic:      "electronics",
				Difficulty: 2,
				Format:     domain.FormatRead,
				URL:        "https://www.explainthatstuff.com/howvalveswork.html",
				Creator:    "Chris Woodford",
				Source:     "explainthatstuff.com",
			},
			{
				ID:         "piece-robot-hands",
				Title:      "Why Robot Hands Are So Hard",
				GapHook:    "Robots can beat you at chess but not at picking up the chess piece. Why?",
				Topic:      "robotics",
				Difficulty: 2,
				Format:     domain.FormatRead,
				URL:        "https://spectrum.ieee.org/robot-hands",
				Creator:    "IEEE Spectrum staff",
				Source:     "spectrum.ieee.org",
			},
			{
				ID:         "piece-longitude",
				Title:      "The Longitude Problem",
				GapHook:    "For 300 years, not knowing where you were east-west killed sailors by the thousand. A carpenter's clock fixed it.",
				Topic:      "history",
				Difficulty: 1,
				Format:     domain.FormatRead,
				URL:        "https://www.rmg.co.uk/stories/topics/longitude-found-john-harrison",
				Creator:    "Royal Museums Greenwich",
				Source:     "rmg.co.uk",
			},
			{
				ID:         "piece-harmonics",
				Title:      "Why Does an Octave Sound 'The Same'?",
				GapHook:    "Two notes, double the frequency, and every culture on earth hears them as one. What is your brain doing?",
				Topic:      "music",
				Difficulty: 2,
				Format:     domain.FormatRead,
				URL:        "https://www.ethanhein.com/wp/2021/octave-equivalence/",
				Creator:    "Ethan Hein",
				Source:     "ethanhein.com",
			},
			{
				ID:         "piece-voyager",
				Title:      "The Farthest Human-Made Object",
				GapHook:    "Voyager 1 runs on 1970s hardware with less memory than your car key. It is still phoning home from interstellar space.",
				Topic:      "space",
				Difficulty: 1,
				Format:     domain.FormatRead,
				URL:        "https://science.nasa.gov/mission/voyager/",
				Creator:    "NASA Voyager team",
				Source:     "science.nasa.gov",
			},
			{
				ID:         "piece-steal-artist",
				Title:      "Steal Like an Artist",
				GapHook:    "Every artist you admire started by copying. Where is the line between stealing and becoming?",
				Topic:      "art",
				Difficulty: 1,
				Format:     domain.FormatRead,
				URL:        "https://austinkleon.com/steal/",
				Creator:    "Austin Kleon",
				Source:     "austinkleon.com",
			},
			{
				ID:         "piece-tree-networks",
				Title:      "How Trees Talk to Each Other",
				GapHook:    "Under every forest floor is a fungal network trading sugar and warnings between trees. A wood-wide web.",
				Topic:      "nature",
				Difficulty: 2,
				Format:     domain.FormatRead,
				URL:        "https://www.ted.com/talks/suzanne_simard_how_trees_talk_to_each_other",
				Creator:    "Suzanne Simard",
				Source:     "ted.com",
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
