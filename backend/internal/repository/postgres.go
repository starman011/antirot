// Package repository owns all data access and returns domain models, never raw rows.
package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/starman011/antirot/backend/internal/domain"
)

func NewPool(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("repository: pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("repository: ping: %w", err)
	}
	return pool, nil
}

type PostgresPieceRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresPieceRepository(pool *pgxpool.Pool) *PostgresPieceRepository {
	return &PostgresPieceRepository{pool: pool}
}

func (r *PostgresPieceRepository) FindByTopics(ctx context.Context, topics []string) ([]domain.Piece, error) {
	const q = `
		SELECT id, title, gap_hook, topic, difficulty, format, url, creator, source
		FROM pieces
		WHERE COALESCE(cardinality($1::text[]), 0) = 0 OR topic = ANY($1::text[])`

	rows, err := r.pool.Query(ctx, q, topics)
	if err != nil {
		return nil, fmt.Errorf("repository: find pieces: %w", err)
	}
	defer rows.Close()

	var out []domain.Piece
	for rows.Next() {
		var p domain.Piece
		if err := rows.Scan(&p.ID, &p.Title, &p.GapHook, &p.Topic, &p.Difficulty, &p.Format, &p.URL, &p.Creator, &p.Source); err != nil {
			return nil, fmt.Errorf("repository: scan piece: %w", err)
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

type PostgresCheckInRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresCheckInRepository(pool *pgxpool.Pool) *PostgresCheckInRepository {
	return &PostgresCheckInRepository{pool: pool}
}

func (r *PostgresCheckInRepository) Save(ctx context.Context, c domain.CheckIn) (domain.CheckIn, error) {
	const q = `INSERT INTO checkins (state) VALUES ($1) RETURNING id, created_at`

	if err := r.pool.QueryRow(ctx, q, string(c.State)).Scan(&c.ID, &c.CreatedAt); err != nil {
		return domain.CheckIn{}, fmt.Errorf("repository: save checkin: %w", err)
	}
	return c, nil
}
