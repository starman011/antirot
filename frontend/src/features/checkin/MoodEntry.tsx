import { useState } from 'react';
import Strands from '../../components/Strands';
import type { CheckInState } from '../../lib/api/client';

// Keyword stub until the AI recommendation service interprets mood text (T302).
// Mood text is sensitive: kept in memory only, never persisted or sent anywhere yet.
function interpretMood(text: string): CheckInState {
  const t = text.toLowerCase();
  if (/scroll|feed|phone|insta|reel|tiktok/.test(t)) return 'doomscrolling';
  if (/tired|unmotivat|lazy|stuck|drained|meh/.test(t)) return 'unmotivated';
  if (/restless|anxious|anxiety|jittery|wired/.test(t)) return 'restless';
  if (/focus|concentrate|work|study|deep/.test(t)) return 'seeking_focus';
  return 'just_curious';
}

interface Props {
  onDone: (state: CheckInState) => void;
  onSkip: () => void;
}

// Skippable, with a "surprise me" path (Principle IV).
export function MoodEntry({ onDone, onSkip }: Props) {
  const [text, setText] = useState('');

  return (
    <section aria-label="Write how you feel">
      <span className="tagline">Welcome back, human</span>
      <h1>How are you, really?</h1>
      <div className="mood-stage">
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
        <textarea
          className="mood-input"
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="Type it out. No one reads this but you."
          rows={4}
        />
      </div>
      <div className="mood-actions">
        <button className="btn-pill" disabled={!text.trim()} onClick={() => onDone(interpretMood(text))}>
          Interpret and curate
        </button>
        <button className="skip-link" onClick={onSkip}>
          Skip, surprise me
        </button>
      </div>
    </section>
  );
}
