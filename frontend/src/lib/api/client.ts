// Typed API client, the only place the frontend talks to the backend. Shapes mirror api/openapi.yaml.

export type CheckInState =
  | 'restless'
  | 'doomscrolling'
  | 'unmotivated'
  | 'seeking_focus'
  | 'just_curious';

export type PieceFormat = 'read' | 'audio';

export interface Piece {
  id: string;
  title: string;
  gap_hook: string;
  topic: string;
  difficulty: number;
  format: PieceFormat;
  url: string;
  creator: string;
  source: string;
}

export interface ApiError {
  code: string;
  message: string;
}

export class ApiRequestError extends Error {
  readonly status: number;
  readonly apiError: ApiError;

  constructor(status: number, apiError: ApiError) {
    super(`${apiError.code}: ${apiError.message}`);
    this.status = status;
    this.apiError = apiError;
  }
}

const BASE = import.meta.env.VITE_API_BASE_URL ?? '/api/v1';

async function request<T>(path: string, body?: object): Promise<T> {
  const res = await fetch(`${BASE}${path}`, {
    method: body ? 'POST' : 'GET',
    headers: { Accept: 'application/json', ...(body && { 'Content-Type': 'application/json' }) },
    ...(body && { body: JSON.stringify(body) }),
  });

  if (!res.ok) {
    const body = (await res.json().catch(() => null)) as { error?: ApiError } | null;
    throw new ApiRequestError(
      res.status,
      body?.error ?? { code: 'unknown', message: res.statusText },
    );
  }
  return (await res.json()) as T;
}

/** State omitted means check-in skipped. */
export function getSessionPiece(state?: CheckInState, interests?: string[]): Promise<Piece> {
  const params = new URLSearchParams();
  if (state) params.set('state', state);
  if (interests?.length) params.set('interests', interests.join(','));
  const qs = params.toString();
  return request<Piece>(`/session/piece${qs ? `?${qs}` : ''}`);
}

export function getHealth(): Promise<{ status: string }> {
  return request<{ status: string }>(`/health`);
}

/** Mood text is interpreted server-side and discarded, never stored. */
export function interpretMood(text: string): Promise<CheckInState> {
  return request<{ state: CheckInState }>(`/interpret`, { text }).then((r) => r.state);
}

export function postCheckIn(state: CheckInState): Promise<void> {
  return request<{ id: string }>(`/checkins`, { state }).then(() => undefined);
}
