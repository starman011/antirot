// Visit-day tracking, localStorage only for now (backend in Phase 2).
// Framing rule (Principle IV): cumulative and forgiving, never "broken".

const KEY = 'antirot.visits';

function load(): string[] {
  try {
    const raw = localStorage.getItem(KEY);
    return raw ? (JSON.parse(raw) as string[]) : [];
  } catch {
    return [];
  }
}

export function recordVisit(): void {
  const today = new Date().toISOString().slice(0, 10);
  const days = load();
  if (!days.includes(today)) {
    days.push(today);
    localStorage.setItem(KEY, JSON.stringify(days));
  }
}

export interface VisitStats {
  totalDays: number;
  thisMonth: number;
  firstVisit: string | null;
}

export function getStats(): VisitStats {
  const days = load();
  const month = new Date().toISOString().slice(0, 7);
  return {
    totalDays: days.length,
    thisMonth: days.filter((d) => d.startsWith(month)).length,
    firstVisit: days[0] ?? null,
  };
}
