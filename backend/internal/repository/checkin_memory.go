package repository

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/starman011/antirot/backend/internal/domain"
)

// MemoryCheckInRepository is the Phase-1 stand-in; Postgres replaces it in T104.
type MemoryCheckInRepository struct {
	mu       sync.Mutex
	checkins []domain.CheckIn
}

func NewMemoryCheckInRepository() *MemoryCheckInRepository {
	return &MemoryCheckInRepository{}
}

func (r *MemoryCheckInRepository) Save(_ context.Context, c domain.CheckIn) (domain.CheckIn, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return domain.CheckIn{}, err
	}
	c.ID = hex.EncodeToString(b)
	c.CreatedAt = time.Now().UTC()

	r.mu.Lock()
	r.checkins = append(r.checkins, c)
	r.mu.Unlock()
	return c, nil
}
