import { getStats } from '../lib/streak';

// Forgiving by design (Principle IV): cumulative growing map,
// no consecutive-day pressure, nothing ever "breaks".
export function Streak() {
  const stats = getStats();

  return (
    <div className="page-inner centered">
      <span className="tagline">Your clarity log</span>
      <h1>{stats.totalDays} {stats.totalDays === 1 ? 'day' : 'days'} of showing up.</h1>
      <p className="dim">Every visit adds a cell. Missing a day changes nothing, this map only grows.</p>

      <div className="stat-row">
        <div className="stat-item">
          <h3>{stats.totalDays}</h3>
          <p>Total days</p>
        </div>
        <div className="stat-item">
          <h3>{stats.thisMonth}</h3>
          <p>This month</p>
        </div>
        <div className="stat-item">
          <h3>{stats.firstVisit ?? 'today'}</h3>
          <p>First visit</p>
        </div>
      </div>

      <div className="streak-grid">
        {Array.from({ length: stats.totalDays }, (_, i) => (
          <div key={i} className="streak-cell" />
        ))}
      </div>
    </div>
  );
}
