package service

import (
	"context"
	"fmt"
	"regexp"

	"github.com/starman011/antirot/backend/internal/domain"
)

type CheckInRepository interface {
	Save(ctx context.Context, c domain.CheckIn) (domain.CheckIn, error)
}

type CheckInService struct {
	checkins CheckInRepository
}

func NewCheckInService(checkins CheckInRepository) *CheckInService {
	return &CheckInService{checkins: checkins}
}

func (s *CheckInService) Record(ctx context.Context, state domain.State) (domain.CheckIn, error) {
	if !state.Valid() {
		return domain.CheckIn{}, fmt.Errorf("checkin: invalid state %q", state)
	}
	saved, err := s.checkins.Save(ctx, domain.CheckIn{State: state})
	if err != nil {
		return domain.CheckIn{}, fmt.Errorf("checkin: save: %w", err)
	}
	return saved, nil
}

var moodPatterns = []struct {
	re    *regexp.Regexp
	state domain.State
}{
	{regexp.MustCompile(`scroll|feed|phone|insta|reel|tiktok`), domain.StateDoomscrolling},
	{regexp.MustCompile(`tired|unmotivat|lazy|stuck|drained|meh`), domain.StateUnmotivated},
	{regexp.MustCompile(`restless|anxious|anxiety|jittery|wired`), domain.StateRestless},
	{regexp.MustCompile(`focus|concentrate|work|study|deep`), domain.StateSeekingFocus},
}

// InterpretMood is the seam for the recommendation agent (T302): keyword
// stub now, ML service later. Text is interpreted and discarded, never stored.
func (s *CheckInService) InterpretMood(text string) domain.State {
	lower := []byte(text)
	for i, c := range lower {
		if c >= 'A' && c <= 'Z' {
			lower[i] = c + 32
		}
	}
	for _, p := range moodPatterns {
		if p.re.Match(lower) {
			return p.state
		}
	}
	return domain.StateJustCurious
}
