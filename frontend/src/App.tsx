import { useEffect } from 'react';
import { BrowserRouter, Route, Routes, useLocation } from 'react-router-dom';
import { NavBar } from './components/NavBar';
import { recordVisit } from './lib/streak';
import { About } from './pages/About';
import { Community } from './pages/Community';
import { Home } from './pages/Home';
import { Random } from './pages/Random';
import { Streak } from './pages/Streak';

function Shell() {
  const bare = useLocation().pathname === '/';

  return (
    <div className={`viewport${bare ? ' bare' : ''}`}>
      <NavBar />
      <main className={`page-container${bare ? ' bare' : ''}`}>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/community" element={<Community />} />
          <Route path="/random" element={<Random />} />
          <Route path="/streak" element={<Streak />} />
          <Route path="/about" element={<About />} />
        </Routes>
      </main>
    </div>
  );
}

export default function App() {
  useEffect(recordVisit, []);

  return (
    <BrowserRouter>
      <Shell />
    </BrowserRouter>
  );
}
