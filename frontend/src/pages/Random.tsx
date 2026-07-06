import { useState } from 'react';
import { getSessionPiece, type Piece } from '../lib/api/client';

const INTERESTS = ['electronics', 'robotics', 'history', 'music', 'space', 'art', 'nature'];

// Static pool until the recommendation agent learns preferences (T302).
const ACTIVITIES = [
  { topic: 'art', text: 'Draw the object nearest to you in exactly five minutes. Bad drawings count double.' },
  { topic: 'electronics', text: 'Open a dead gadget and identify three components before looking anything up.' },
  { topic: 'robotics', text: 'Sketch a robot that solves your most boring chore. Label its parts.' },
  { topic: 'music', text: 'Make a rhythm using only objects on your desk. Record 30 seconds.' },
  { topic: 'history', text: 'Pick a building you pass daily and find out what stood there 100 years ago.' },
  { topic: 'space', text: 'Step outside tonight and find one constellation without an app.' },
  { topic: 'nature', text: 'Walk ten minutes and collect three textures. Describe them in writing.' },
];

type Result = { kind: 'piece'; piece: Piece } | { kind: 'activity'; text: string; topic: string };

export function Random() {
  const [selected, setSelected] = useState<string[]>([]);
  const [result, setResult] = useState<Result | null>(null);
  const [error, setError] = useState<string | null>(null);

  const toggle = (i: string) =>
    setSelected((s) => (s.includes(i) ? s.filter((x) => x !== i) : [...s, i]));

  const recommendRead = () => {
    setError(null);
    getSessionPiece('just_curious', selected)
      .then((piece) => setResult({ kind: 'piece', piece }))
      .catch((e: unknown) => setError(e instanceof Error ? e.message : 'request failed'));
  };

  const recommendActivity = () => {
    setError(null);
    const pool = selected.length ? ACTIVITIES.filter((a) => selected.includes(a.topic)) : ACTIVITIES;
    const pick = pool[Math.floor(Math.random() * pool.length)] ?? ACTIVITIES[0]!;
    setResult({ kind: 'activity', text: pick.text, topic: pick.topic });
  };

  return (
    <div className="page-inner">
      <span className="tagline">Agent recommended</span>
      <h1>Random.</h1>
      <p className="dim">Pick what pulls you. The agent learns from every reaction and suggests one blog to read or one thing to make. One at a time, always finite.</p>

      <div className="chip-row">
        {INTERESTS.map((i) => (
          <button
            key={i}
            className={`chip${selected.includes(i) ? ' selected' : ''}`}
            onClick={() => toggle(i)}
          >
            {i}
          </button>
        ))}
      </div>

      <div className="session-actions">
        <button className="btn-pill" onClick={recommendRead}>
          Read something
        </button>
        <button className="btn-pill" onClick={recommendActivity}>
          Make something
        </button>
      </div>

      {result?.kind === 'piece' && (
        <div className="result-card">
          <span className="provenance">{result.piece.topic} · by {result.piece.creator} · {result.piece.source}</span>
          <h2 className="result-title">{result.piece.title}</h2>
          <p className="dim">{result.piece.gap_hook}</p>
          <a className="btn-pill" href={result.piece.url} target="_blank" rel="noreferrer">
            {result.piece.format === 'audio' ? 'Listen' : 'Read'}
          </a>
        </div>
      )}
      {result?.kind === 'activity' && (
        <div className="result-card">
          <span className="provenance">offline · {result.topic}</span>
          <p className="dim">{result.text}</p>
        </div>
      )}
      {error && <p role="alert">{error}</p>}
    </div>
  );
}
