import { useState } from 'react';
import { CheckIn } from './features/checkin/CheckIn';
import { MoodEntry } from './features/checkin/MoodEntry';
import { getSessionPiece, type CheckInState, type Piece } from './lib/api/client';

export default function App() {
  const [piece, setPiece] = useState<Piece | null>(null);
  const [writing, setWriting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const startSession = (state?: CheckInState) => {
    setError(null);
    setWriting(false);
    getSessionPiece(state)
      .then(setPiece)
      .catch((e: unknown) => setError(e instanceof Error ? e.message : 'request failed'));
  };

  return (
    <div className="viewport">
      <main className="page-container">
        <div className="split-layout">
          <aside className="visual-soul" style={{ backgroundImage: "url('/img/tech.jpg')" }}>
            <span className="tagline">
              {piece ? `Daily geek-out // ${piece.topic}` : 'The Sanctuary'}
            </span>
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
            ) : writing ? (
              <MoodEntry onDone={startSession} onBack={() => setWriting(false)} />
            ) : (
              <CheckIn onSelect={startSession} onWrite={() => setWriting(true)} />
            )}
            {error && <p role="alert">{error}</p>}
          </section>
        </div>
      </main>
    </div>
  );
}
