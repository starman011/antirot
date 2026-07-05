package controller

import (
	"errors"
	"net/http"
	"strings"

	"github.com/starman011/antirot/backend/internal/domain"
	"github.com/starman011/antirot/backend/internal/service"
)

// Provenance fields are always present (Principle II).
type pieceResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	GapHook    string `json:"gap_hook"`
	Topic      string `json:"topic"`
	Difficulty int    `json:"difficulty"`
	Format     string `json:"format"`
	URL        string `json:"url"`
	Creator    string `json:"creator"`
	Source     string `json:"source"`
}

type sessionController struct {
	sessions *service.SessionService
}

func (c *sessionController) handlePiece(w http.ResponseWriter, r *http.Request) {
	state := domain.State(r.URL.Query().Get("state"))
	if state == "" {
		// Check-in is skippable (Principle IV).
		state = domain.StateJustCurious
	}
	if !state.Valid() {
		writeError(w, http.StatusBadRequest, "invalid_state", "unknown check-in state")
		return
	}

	var interests []string
	if raw := r.URL.Query().Get("interests"); raw != "" {
		for _, s := range strings.Split(raw, ",") {
			if s = strings.TrimSpace(s); s != "" {
				interests = append(interests, s)
			}
		}
	}

	piece, err := c.sessions.PickPiece(r.Context(), state, interests)
	switch {
	case errors.Is(err, service.ErrNoPiece):
		writeError(w, http.StatusNotFound, "no_piece", "nothing curated for this yet, try broader interests")
		return
	case err != nil:
		writeError(w, http.StatusInternalServerError, "internal", "something went wrong")
		return
	}

	writeJSON(w, http.StatusOK, pieceResponse{
		ID:         piece.ID,
		Title:      piece.Title,
		GapHook:    piece.GapHook,
		Topic:      piece.Topic,
		Difficulty: piece.Difficulty,
		Format:     string(piece.Format),
		URL:        piece.URL,
		Creator:    piece.Creator,
		Source:     piece.Source,
	})
}
