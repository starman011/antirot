package controller

import (
	"encoding/json"
	"net/http"

	"github.com/starman011/antirot/backend/internal/domain"
	"github.com/starman011/antirot/backend/internal/service"
)

const maxBodyBytes = 4 << 10

type checkinController struct {
	checkins *service.CheckInService
}

func (c *checkinController) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		State string `json:"state"`
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_body", "expected json with state")
		return
	}

	saved, err := c.checkins.Record(r.Context(), domain.State(req.State))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid_state", "unknown check-in state")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]string{
		"id":         saved.ID,
		"state":      string(saved.State),
		"created_at": saved.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func (c *checkinController) handleInterpret(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text string `json:"text"`
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Text == "" {
		writeError(w, http.StatusBadRequest, "invalid_body", "expected json with non-empty text")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"state": string(c.checkins.InterpretMood(req.Text)),
	})
}
