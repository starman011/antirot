package service

import (
	"context"
	"testing"

	"github.com/starman011/antirot/backend/internal/domain"
)

type stubCheckInRepo struct {
	saved []domain.CheckIn
}

func (s *stubCheckInRepo) Save(_ context.Context, c domain.CheckIn) (domain.CheckIn, error) {
	c.ID = "c1"
	s.saved = append(s.saved, c)
	return c, nil
}

func TestRecordValidState(t *testing.T) {
	repo := &stubCheckInRepo{}
	svc := NewCheckInService(repo)

	got, err := svc.Record(context.Background(), domain.StateRestless)
	if err != nil {
		t.Fatalf("Record() error = %v", err)
	}
	if got.ID != "c1" || len(repo.saved) != 1 {
		t.Errorf("Record() not persisted: %+v", got)
	}
}

func TestRecordInvalidState(t *testing.T) {
	svc := NewCheckInService(&stubCheckInRepo{})
	if _, err := svc.Record(context.Background(), domain.State("angry")); err == nil {
		t.Error("Record() with invalid state: want error, got nil")
	}
}

func TestInterpretMood(t *testing.T) {
	svc := NewCheckInService(&stubCheckInRepo{})
	cases := map[string]domain.State{
		"been scrolling Reels for an hour": domain.StateDoomscrolling,
		"so TIRED and stuck":               domain.StateUnmotivated,
		"feeling anxious and wired":        domain.StateRestless,
		"need to focus on study":           domain.StateSeekingFocus,
		"hello world":                      domain.StateJustCurious,
	}
	for text, want := range cases {
		if got := svc.InterpretMood(text); got != want {
			t.Errorf("InterpretMood(%q) = %q, want %q", text, got, want)
		}
	}
}
