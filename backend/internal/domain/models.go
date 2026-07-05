// Package domain holds core business models with no transport or storage dependencies.
package domain

import "time"

// State is the user's self-reported internal state at check-in.
type State string

const (
	StateRestless      State = "restless"
	StateDoomscrolling State = "doomscrolling"
	StateUnmotivated   State = "unmotivated"
	StateSeekingFocus  State = "seeking_focus"
	StateJustCurious   State = "just_curious"
)

func (s State) Valid() bool {
	switch s {
	case StateRestless, StateDoomscrolling, StateUnmotivated, StateSeekingFocus, StateJustCurious:
		return true
	}
	return false
}

// Format is how a piece is consumed.
type Format string

const (
	FormatRead  Format = "read"
	FormatAudio Format = "audio"
)

// Piece is one curated, finite, human-made item. Creator and Source are
// required provenance and never empty (Principle II).
type Piece struct {
	ID         string
	Title      string
	GapHook    string
	Topic      string
	Difficulty int // 1 beginner .. 5 deep expertise
	Format     Format
	URL        string
	Creator    string
	Source     string
}

type CheckIn struct {
	ID        string
	State     State
	CreatedAt time.Time
}
