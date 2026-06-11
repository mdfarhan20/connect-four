import "./GameScreen.css";

type Cell = "empty" | "red" | "yellow";

const board: Cell[][] = [
  ["empty", "empty", "empty", "empty", "empty", "empty", "empty"],
  ["empty", "empty", "empty", "empty", "empty", "empty", "empty"],
  ["empty", "empty", "empty", "empty", "empty", "empty", "empty"],
  ["empty", "empty", "empty", "yellow", "empty", "empty", "empty"],
  ["empty", "empty", "red", "red", "yellow", "empty", "empty"],
  ["empty", "red", "yellow", "red", "yellow", "empty", "empty"],
];

function GameScreen() {
  return (
    <section className="game-screen" aria-label="Connect Four game board">
      <aside className="game-screen__side game-screen__side--player" aria-label="Player status">
        <div className="game-player game-player--active">
          <span className="game-player__chip game-player__chip--red" aria-hidden="true" />
          <span className="label-sm">You</span>
          <strong>Terracotta</strong>
          <span className="game-player__score">2</span>
        </div>
      </aside>

      <div className="game-screen__center">
        <div className="game-screen__heading">
          <span className="status-badge status-badge-turn-red">Your turn</span>
          <p className="body-md">Drop a chip into any open column.</p>
        </div>

        <div className="game-board-shell" aria-label="Game board">
          <div className="game-board__columns" aria-hidden="true">
            {Array.from({ length: 7 }).map((_, index) => (
              <button className="game-board__column" type="button" key={index}>
                {index + 1}
              </button>
            ))}
          </div>

          <div className="game-board" role="grid" aria-rowcount={6} aria-colcount={7}>
            {board.flatMap((row, rowIndex) =>
              row.map((cell, colIndex) => (
                <span
                  className={`game-board__cell ${cell !== "empty" ? `game-board__cell--${cell}` : ""}`}
                  role="gridcell"
                  aria-label={
                    cell === "empty"
                      ? `Empty slot row ${rowIndex + 1} column ${colIndex + 1}`
                      : `${cell} chip row ${rowIndex + 1} column ${colIndex + 1}`
                  }
                  key={`${rowIndex}-${colIndex}`}
                />
              )),
            )}
          </div>
        </div>
      </div>

      <aside className="game-screen__side game-screen__side--bot" aria-label="Bot status">
        <div className="game-player">
          <span className="game-player__chip game-player__chip--yellow" aria-hidden="true" />
          <span className="label-sm">Bot</span>
          <strong>Sage</strong>
          <span className="game-player__score">1</span>
        </div>
      </aside>
    </section>
  );
}

export default GameScreen;
