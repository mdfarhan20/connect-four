import type { Board, Cell } from "../../types/game";
import { HEIGHT, WIDTH } from "../../constants/game";
import "./GameScreen.css";

const CELL_LABEL: Record<Cell, string> = {
  0: "empty",
  1: "red",
  2: "yellow",
};

const CELL_NAME: Record<Cell, string> = {
  0: "Empty",
  1: "Red",
  2: "Yellow",
};

type GameScreenProps = {
  board: Board;
  moveLoading: boolean;
  botThinking: boolean;
  onPlayerMove: (col: number) => void;
};

function GameScreen({
  board,
  moveLoading,
  botThinking,
  onPlayerMove,
}: GameScreenProps) {
  return (
    <section className="game-screen" aria-label="Connect Four game board">
      <aside
        className="game-screen__side game-screen__side--player"
        aria-label="Player status"
      >
        <div className="game-player game-player--active">
          <span
            className="game-player__chip game-player__chip--red"
            aria-hidden="true"
          />
          <span className="label-sm">You</span>
          <strong>Terracotta</strong>
          <span className="game-player__score">2</span>
        </div>
      </aside>

      <div className="game-screen__center">
        <div className="game-screen__heading">
          <span className="status-badge status-badge-turn-red">
            {botThinking ? "Bot thinking..." : "Your turn"}
          </span>
          <p className="body-md">
            {botThinking
              ? "Wait till the bot makes a move."
              : "Drop a chip into any open column."}
          </p>
        </div>

        <div className="game-board-shell" aria-label="Game board">
          <div className="game-board__columns" aria-hidden="true">
            {Array.from({ length: WIDTH }).map((_, index) => (
              <button
                className="game-board__column"
                type="button"
                key={index}
                disabled={moveLoading || botThinking}
                onClick={() => onPlayerMove(index)}
              >
                {index + 1}
              </button>
            ))}
          </div>

          <div
            className="game-board"
            role="grid"
            aria-rowcount={HEIGHT}
            aria-colcount={WIDTH}
          >
            {board.flatMap((row: Cell[], rowIndex: number) =>
              row.map((cell: Cell, colIndex: number) => {
                const label = CELL_LABEL[cell];
                const name = CELL_NAME[cell];
                return (
                  <span
                    className={`game-board__cell ${cell !== 0 ? `game-board__cell--${label}` : ""}`}
                    role="gridcell"
                    aria-label={
                      cell === 0
                        ? `Empty slot row ${rowIndex + 1} column ${colIndex + 1}`
                        : `${name} chip row ${rowIndex + 1} column ${colIndex + 1}`
                    }
                    key={`${rowIndex}-${colIndex}`}
                  />
                );
              }),
            )}
          </div>
        </div>
      </div>

      <aside
        className="game-screen__side game-screen__side--bot"
        aria-label="Bot status"
      >
        <div className="game-player">
          <span
            className="game-player__chip game-player__chip--yellow"
            aria-hidden="true"
          />
          <span className="label-sm">Bot</span>
          <strong>Sage</strong>
          <span className="game-player__score">1</span>
        </div>
      </aside>
    </section>
  );
}

export default GameScreen;
