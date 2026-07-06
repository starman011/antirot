import { NavLink } from 'react-router-dom';

const LINKS = [
  { to: '/', label: 'Home' },
  { to: '/community', label: 'Community' },
  { to: '/random', label: 'Random' },
  { to: '/streak', label: 'Streak' },
  { to: '/about', label: 'About' },
];

export function NavBar() {
  return (
    <nav className="top-nav">
      {LINKS.map((l) => (
        <NavLink
          key={l.to}
          to={l.to}
          className={({ isActive }) => `nav-link${isActive ? ' active' : ''}`}
        >
          {l.label}
        </NavLink>
      ))}
    </nav>
  );
}
