import "./HowToPlay.css";

type HowToPlayProps = {
  onBack: () => void;
};

function HowToPlay({ onBack }: HowToPlayProps) {
  return (
    <section className="how-to-play" aria-label="How to play">
      <div className="how-to-play__header">
        <button className="how-to-play__back" type="button" onClick={onBack} aria-label="Back to game">
          <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M19 12H5" />
            <polyline points="12 19 5 12 12 5" />
          </svg>
        </button>
        <h1 className="how-to-play__title">Connect Four</h1>
      </div>

      <div className="how-to-play__hero">
        <span className="how-to-play__badge">Master the game</span>
        <h2 className="how-to-play__heading">How to Play</h2>
        <p className="how-to-play__subtitle">
          Connect Four is a classic two-player strategy game. Simple to learn, challenging to master. Here is everything you need to know to win.
        </p>
      </div>

      <div className="how-to-play__grid">
        <div className="how-to-play__card how-to-play__card--goal">
          <div className="how-to-play__card-icon">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M4 15s1-1 4-1 5 2 8 2 4-1 4-1V3s-1 1-4 1-5-2-8-2-4 1-4 1z" />
              <line x1="4" y1="22" x2="4" y2="15" />
            </svg>
          </div>
          <h3>The Goal</h3>
          <p>
            Be the first player to connect <strong>four</strong> of your colored chips in a row—either horizontally, vertically, or diagonally.
          </p>
          <div className="how-to-play__chip-row">
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
          </div>
        </div>

        <div className="how-to-play__card how-to-play__card--setup">
          <div className="how-to-play__card-icon">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <rect x="3" y="3" width="18" height="18" rx="2" ry="2" />
              <line x1="3" y1="9" x2="21" y2="9" />
              <line x1="9" y1="21" x2="9" y2="9" />
            </svg>
          </div>
          <h3>Game Setup</h3>
          <p>
            The game is played on a vertical board with 7 columns and 6 rows. Players choose a color (usually Terracotta or Sage).
          </p>
          <div className="how-to-play__board-mini">
            {Array.from({ length: 7 }).map((_, i) => (
              <span key={i} className="how-to-play__slot" />
            ))}
          </div>
        </div>

        <div className="how-to-play__card how-to-play__card--turns">
          <div className="how-to-play__card-icon">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <polyline points="17 1 21 5 17 9" />
              <path d="M3 11V9a4 4 0 0 1 4-4h14" />
              <polyline points="7 23 3 19 7 15" />
              <path d="M21 13v2a4 4 0 0 1-4 4H3" />
            </svg>
          </div>
          <h3>Taking Turns</h3>
          <ol className="how-to-play__steps">
            <li>
              <span className="how-to-play__step-num">1</span>
              <span>Players alternate turns dropping one chip into any of the 7 columns.</span>
            </li>
            <li>
              <span className="how-to-play__step-num">2</span>
              <span>The chip falls to the lowest available space in that column.</span>
            </li>
            <li>
              <span className="how-to-play__step-num">3</span>
              <span>You cannot move a chip once it has been dropped.</span>
            </li>
          </ol>
        </div>

        <div className="how-to-play__card how-to-play__card--tips">
          <div className="how-to-play__card-icon">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
              <path d="M9 18h6" />
              <path d="M10 22h4" />
              <path d="M15.09 14c.18-.98.65-1.74 1.41-2.5A4.65 4.65 0 0 0 18 8 6 6 0 0 0 6 8c0 1 .23 2.23 1.5 3.5A4.61 4.61 0 0 1 8.91 14" />
            </svg>
          </div>
          <h3>Pro Tips</h3>
          <div className="how-to-play__tips-list">
            <div className="how-to-play__tip">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <path d="M12 2a4 4 0 0 0-4 4c0 2 2 4 4 4s4-2 4-4a4 4 0 0 0-4-4z" />
                <path d="M12 14c-4 0-6 2-6 4" />
              </svg>
              <span>Control the center column early.</span>
            </div>
            <div className="how-to-play__tip">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z" />
                <circle cx="12" cy="12" r="3" />
              </svg>
              <span>Watch your opponent's moves to block threats.</span>
            </div>
            <div className="how-to-play__tip">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                <line x1="6" y1="3" x2="6" y2="15" />
                <circle cx="18" cy="6" r="3" />
                <circle cx="6" cy="18" r="3" />
                <path d="M18 9a9 9 0 0 1-9 9" />
              </svg>
              <span>Create 'forks' where you have two ways to win.</span>
            </div>
          </div>
        </div>

        <div className="how-to-play__card how-to-play__card--pattern">
          <span className="how-to-play__pattern-label">Pattern 01</span>
          <div className="how-to-play__pattern-visual">
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
          </div>
          <h4>Horizontal Win</h4>
          <p>Four chips in the same row.</p>
        </div>

        <div className="how-to-play__card how-to-play__card--pattern">
          <span className="how-to-play__pattern-label">Pattern 02</span>
          <div className="how-to-play__pattern-visual how-to-play__pattern-visual--vertical">
            <span className="how-to-play__chip how-to-play__chip--teal" />
            <span className="how-to-play__chip how-to-play__chip--teal" />
            <span className="how-to-play__chip how-to-play__chip--teal" />
            <span className="how-to-play__chip how-to-play__chip--teal" />
          </div>
          <h4>Vertical Win</h4>
          <p>Four chips in the same column.</p>
        </div>

        <div className="how-to-play__card how-to-play__card--pattern">
          <span className="how-to-play__pattern-label">Pattern 03</span>
          <div className="how-to-play__pattern-visual how-to-play__pattern-visual--diagonal">
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
            <span className="how-to-play__chip how-to-play__chip--red" />
          </div>
          <h4>Diagonal Win</h4>
          <p>Four chips slanted across the grid.</p>
        </div>
      </div>

      <div className="how-to-play__cta">
        <div className="how-to-play__cta-text">
          <h3>Ready to challenge?</h3>
          <p>Start a new game with a friend or the CPU.</p>
        </div>
        <div className="how-to-play__cta-actions">
          <button className="btn btn-primary" type="button" onClick={onBack}>Play the Game</button>
        </div>
      </div>
    </section>
  );
}

export default HowToPlay;
