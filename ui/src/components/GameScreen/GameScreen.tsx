import type { Cell, BoardState, Player } from "../../types/game";
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
  gameState: BoardState;
  moveLoading: boolean;
  botThinking: boolean;
  onPlayerMove: (col: number) => void;
  onResetGame: () => void;
  playerColor: Player;
};

function GameScreen({
  gameState,
  moveLoading,
  botThinking,
  onPlayerMove,
  onResetGame,
  playerColor,
}: GameScreenProps) {
  const { board, winner, player, isDraw } = gameState;
  const isGameOver = winner !== 0 || isDraw;

  const playerWon = winner === playerColor;

  let statusTitle: string;
  let statusDescription: string;

  if (isGameOver) {
    if (isDraw) {
      statusTitle = "It's a Draw!";
      statusDescription = "No one wins this round.";
    } else if (playerWon) {
      statusTitle = "You Win!";
      statusDescription = "You outplayed the bot.";
    } else {
      statusTitle = "Bot Wins!";
      statusDescription = "Better luck next time.";
    }
  } else if (botThinking) {
    statusTitle = "Bot Thinking...";
    statusDescription = "Wait for the bot to make a move.";
  } else {
    statusTitle = "Your Turn";
    statusDescription = "Drop a chip into any open column.";
  }

  return (
    <section className="game-screen" aria-label="Connect Four game board">
      <div className="game-screen__board">
        <div className="game-board-shell" aria-label="Game board">
          <div className="game-board__columns" aria-hidden="true">
            {Array.from({ length: WIDTH }).map((_, index) => (
              <button
                className="game-board__column"
                type="button"
                key={index}
                disabled={moveLoading || botThinking || isGameOver}
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

      <aside className="game-screen__status" aria-label="Game status">
        <div className={`status-card ${isGameOver ? "status-card--over" : botThinking ? "status-card--thinking" : "status-card--active"}`}>
          <div className="status-card__indicator">
            {isGameOver ? (
              isDraw ? (
                <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
                  <circle cx="12" cy="12" r="10" />
                  <line x1="12" y1="8" x2="12" y2="12" />
                  <line x1="12" y1="16" x2="12.01" y2="16" />
                </svg>
              ) : (
                <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round">
                  <path d="M6 9H4.5a2.5 2.5 0 0 1 0-5C7 4 9 6 9 9v.5A2.5 2.5 0 0 1 6.5 12H6" />
                  <path d="M18 9h1.5a2.5 2.5 0 0 0 0-5C17 4 15 6 15 9v.5A2.5 2.5 0 0 0 17.5 12H18" />
                  <path d="M4 22h16" />
                  <path d="M10 22V2h4v20" />
                </svg>
              )
            ) : (
              <span className={`status-card__dot ${player === 1 ? "status-card__dot--red" : "status-card__dot--yellow"}`}>
                {botThinking && <span className="status-card__pulse" />}
              </span>
            )}
          </div>

          <h3 className="status-card__title">{statusTitle}</h3>
          <p className="status-card__desc">{statusDescription}</p>

          {isGameOver && (
            <button className="status-card__btn" type="button" onClick={onResetGame}>
              Play Again
            </button>
          )}
        </div>
      </aside>
    </section>
  );
}

export default GameScreen;
