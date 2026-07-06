import { useState } from 'react';
import Strands from '../components/Strands';
import {
  getSessionPiece,
  interpretMood,
  postCheckIn,
  type CheckInState,
  type Piece,
} from '../lib/api/client';

export function Home() {
  const [text, setText] = useState('');
  const [piece, setPiece] = useState<Piece | null>(null);
  const [error, setError] = useState<string | null>(null);

  const fail = (e: unknown) => setError(e instanceof Error ? e.message : 'request failed');

  const startSession = (state?: CheckInState) => {
    setError(null);
    if (state) postCheckIn(state).catch(() => undefined);
    getSessionPiece(state).then(setPiece).catch(fail);
  };

  const interpretAndStart = () => {
    setError(null);
    interpretMood(text).then(startSession).catch(fail);
  };

  if (piece) {
    return (
      <div className="home-stage">
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
        {error && <p role="alert">{error}</p>}
      </div>
    );
  }

  return (
    <div className="home-stage">
      <div className="orb-stage">
        <Strands
          colors={['#5d1b8d', '#88cc20', '#0064f2']}
          count={3}
          speed={0.4}
          amplitude={1.7}
          waviness={1.6}
          thickness={0.5}
          glow={0.7}
          taper={2.8}
          spread={0.5}
          intensity={0.6}
          saturation={2}
          opacity={1}
          scale={1.5}
          glass
          refraction={3}
          dispersion={1.5}
          glassSize={1}
          hueShift={0}
        />
      </div>
      <input
        className="feel-bar"
        type="text"
        value={text}
        onChange={(e) => setText(e.target.value)}
        onKeyDown={(e) => {
          if (e.key === 'Enter' && text.trim()) interpretAndStart();
        }}
        placeholder="What do you feel?"
        aria-label="What do you feel?"
      />
      <button className="skip-link" onClick={() => startSession(undefined)}>
        Surprise me
      </button>
      {error && <p role="alert">{error}</p>}
    </div>
  );
}
