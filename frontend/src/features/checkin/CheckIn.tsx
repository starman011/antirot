import type { CheckInState } from '../../lib/api/client';

const STATES: { value: CheckInState; label: string }[] = [
  { value: 'restless', label: 'Restless' },
  { value: 'doomscrolling', label: 'Doomscrolling' },
  { value: 'unmotivated', label: 'Unmotivated' },
  { value: 'seeking_focus', label: 'Seeking focus' },
  { value: 'just_curious', label: 'Just curious' },
];

interface Props {
  onSelect: (state?: CheckInState) => void;
  onWrite: () => void;
}

// Skippable, with a "surprise me" path (Principle IV).
export function CheckIn({ onSelect, onWrite }: Props) {
  return (
    <section aria-label="Spirit check">
      <span className="tagline">Welcome back, human</span>
      <h1>Spirit Check.</h1>
      <p className="dim">How is your internal state today? We'll curate your session accordingly.</p>
      <div className="check-in-grid">
        {STATES.map((s) => (
          <button key={s.value} className="check-btn" onClick={() => onSelect(s.value)}>
            {s.label}
          </button>
        ))}
        <button className="check-btn" onClick={() => onSelect(undefined)}>
          Surprise me
        </button>
        <button className="check-btn" onClick={onWrite}>
          Write it out
        </button>
      </div>
      <button className="skip-link" onClick={() => onSelect(undefined)}>
        Skip check-in
      </button>
    </section>
  );
}
