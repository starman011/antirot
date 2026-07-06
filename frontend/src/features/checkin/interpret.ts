import type { CheckInState } from '../../lib/api/client';

// Keyword stub until the AI recommendation service interprets mood text (T302).
// Mood text is sensitive: kept in memory only, never persisted or sent anywhere yet.
export function interpretMood(text: string): CheckInState {
  const t = text.toLowerCase();
  if (/scroll|feed|phone|insta|reel|tiktok/.test(t)) return 'doomscrolling';
  if (/tired|unmotivat|lazy|stuck|drained|meh/.test(t)) return 'unmotivated';
  if (/restless|anxious|anxiety|jittery|wired/.test(t)) return 'restless';
  if (/focus|concentrate|work|study|deep/.test(t)) return 'seeking_focus';
  return 'just_curious';
}
