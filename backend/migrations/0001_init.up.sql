-- Initial schema. Run with golang-migrate:
--   migrate -path backend/migrations -database "$DATABASE_URL" up

CREATE TABLE users (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email       TEXT UNIQUE NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Interests: chip-based topics a user follows.
CREATE TABLE interests (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    topic       TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (user_id, topic)
);

-- Curated, human-made pieces. Provenance columns are NOT NULL by design
-- (constitution Principle II: human-made content only, always surfaced).
CREATE TABLE pieces (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title       TEXT NOT NULL,
    gap_hook    TEXT NOT NULL,
    topic       TEXT NOT NULL,
    difficulty  SMALLINT NOT NULL CHECK (difficulty BETWEEN 1 AND 5),
    format      TEXT NOT NULL CHECK (format IN ('read', 'audio')),
    url         TEXT NOT NULL,
    creator     TEXT NOT NULL CHECK (creator <> ''),
    source      TEXT NOT NULL CHECK (source <> ''),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX pieces_topic_idx ON pieces (topic);

-- Check-ins: sensitive emotional-state data (constitution: Security).
-- Minimal columns only; user-deletable via ON DELETE CASCADE.
CREATE TABLE checkins (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    state       TEXT NOT NULL CHECK (
        state IN ('restless', 'doomscrolling', 'unmotivated', 'seeking_focus', 'just_curious')
    ),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX checkins_user_idx ON checkins (user_id, created_at);
