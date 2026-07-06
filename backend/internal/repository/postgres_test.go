package repository

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/starman011/antirot/backend/internal/domain"
)

// Integration tests; run with TEST_DATABASE_URL pointing at a migrated antirot DB.
func testPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}
	pool, err := NewPool(context.Background(), url)
	if err != nil {
		t.Fatalf("NewPool() error = %v", err)
	}
	t.Cleanup(pool.Close)
	return pool
}

func TestPostgresPieces(t *testing.T) {
	repo := NewPostgresPieceRepository(testPool(t))

	all, err := repo.FindByTopics(context.Background(), nil)
	if err != nil {
		t.Fatalf("FindByTopics(nil) error = %v", err)
	}
	if len(all) == 0 {
		t.Fatal("FindByTopics(nil) returned no seeded pieces")
	}
	for _, p := range all {
		if p.Creator == "" || p.Source == "" {
			t.Errorf("piece %q missing provenance", p.ID)
		}
	}

	nature, err := repo.FindByTopics(context.Background(), []string{"nature"})
	if err != nil {
		t.Fatalf("FindByTopics(nature) error = %v", err)
	}
	for _, p := range nature {
		if p.Topic != "nature" {
			t.Errorf("topic filter leaked: got %q", p.Topic)
		}
	}
}

func TestPostgresCheckIns(t *testing.T) {
	repo := NewPostgresCheckInRepository(testPool(t))

	saved, err := repo.Save(context.Background(), domain.CheckIn{State: domain.StateJustCurious})
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}
	if saved.ID == "" || saved.CreatedAt.IsZero() {
		t.Errorf("Save() incomplete: %+v", saved)
	}
}
