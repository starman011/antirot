import { useState } from 'react';
import { MoodEntry } from '../features/checkin/MoodEntry';
import { getSessionPiece, type CheckInState, type Piece } from '../lib/api/client';

export function Home() {
  const [piece, setPiece] = useState<Piece | null>(null);
  const [error, setError] = useState<string | null>(null);

  const startSession = (state?: CheckInState) => {
    setError(null);
    getSessionPiece(state)
      .then(setPiece)
      .catch((e: unknown) => setError(e instanceof Error ? e.message : 'request failed'));
  };

  return (
    <div className="split-layout">
      <aside className="visual-soul" style={{ backgroundImage: "url('/img/tech.jpg')" }}>
        <span className="tagline">{piece ? `Daily geek-out // ${piece.topic}` : 'The Sanctuary'}</span>
        <h2>{piece ? piece.title : 'Step into the silence.'}</h2>
      </aside>

      <section className="flow-soul">
        {piece ? (
          <>
            <span className="tagline">{piece.topic} · Human curated</span>
            <h1>{piece.title}</h1>
            <p className="dim">{piece.gap_hook}</p>
            <p className="provenance">
              By {piece.creator} · {piece.source}
            </p>
            <div className="session-actions">
              <a className="btn-pill" href={piece.url} target="_blank" rel="noreferrer">
                {piece.format === 'audio' ? 'Listen' : 'Read'}
              </a>
              <button className="skip-link" onClick={() => setPiece(null)}>
                Complete session
              </button>
            </div>
          </>
        ) : (
          <MoodEntry onDone={startSession} onSkip={() => startSession(undefined)} />
        )}
        {error && <p role="alert">{error}</p>}
      </section>
    </div>
  );
}
