import { useState } from 'react';
import { CheckIn } from './features/checkin/CheckIn';
import { getSessionPiece, type CheckInState, type Piece } from './lib/api/client';

export default function App() {
  const [piece, setPiece] = useState<Piece | null>(null);
  const [error, setError] = useState<string | null>(null);

  const startSession = (state?: CheckInState) => {
    setError(null);
    getSessionPiece(state)
      .then(setPiece)
      .catch((e: unknown) => setError(e instanceof Error ? e.message : 'request failed'));
  };

  return (
    <main className="glass-panel">
      {piece ? (
        <section aria-label="Your piece">
          <span className="tagline">Daily geek-out // {piece.topic}</span>
          <h1>{piece.title}</h1>
          <p className="dim">{piece.gap_hook}</p>
          <p className="provenance">
            By {piece.creator} · {piece.source}
          </p>
          <a className="check-btn" href={piece.url} target="_blank" rel="noreferrer">
            {piece.format === 'audio' ? 'Listen' : 'Read'}
          </a>
          <button className="skip-link" onClick={() => setPiece(null)}>
            That's my session, done for today
          </button>
        </section>
      ) : (
        <CheckIn onSelect={startSession} />
      )}
      {error && <p role="alert">{error}</p>}
    </main>
  );
}
