import { useState } from 'react';

interface Post {
  author: string;
  text: string;
  topic: string;
}

// Local-only stub; backend threads land in Phase 3 (T304).
const SEED: Post[] = [
  { author: 'mira', topic: 'electronics', text: 'Finished the valve amplifier read from last week and built a one-tube preamp. Happy to share schematics.' },
  { author: 'jon', topic: 'space', text: 'The piece on slow interplanetary trajectories sent me down a wonderful two-evening rabbit hole. Anyone want a reading buddy?' },
];

export function Community() {
  const [posts, setPosts] = useState<Post[]>(SEED);
  const [draft, setDraft] = useState('');

  const submit = () => {
    if (!draft.trim()) return;
    setPosts([{ author: 'you', topic: 'general', text: draft.trim() }, ...posts]);
    setDraft('');
  };

  return (
    <div className="page-inner">
      <span className="tagline">Humans talking to humans</span>
      <h1>Community.</h1>
      <p className="dim">Discuss what you are reading, making, learning. No likes, no follower counts, no ranking. Just conversation.</p>

      <div className="composer">
        <textarea
          className="mood-input composer-input"
          value={draft}
          onChange={(e) => setDraft(e.target.value)}
          placeholder="Share something worth discussing"
          rows={3}
        />
        <button className="btn-pill" disabled={!draft.trim()} onClick={submit}>
          Post
        </button>
      </div>

      <ul className="post-list">
        {posts.map((p, i) => (
          <li key={i} className="post">
            <span className="provenance">{p.author} · {p.topic}</span>
            <p>{p.text}</p>
          </li>
        ))}
      </ul>
    </div>
  );
}
