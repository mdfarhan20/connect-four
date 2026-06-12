import "./HomeScreen.css";
import type { Player } from "../../types/game";

type PlayerColor = "red" | "yellow";

const colorOptions: Array<{
  color: PlayerColor;
  label: string;
  description: string;
  value: Player;
}> = [
  {
    color: "red",
    label: "Terracotta",
    value: 1,
    description: "Play first with the warm red chips.",
  },
  {
    color: "yellow",
    label: "Sage",
    value: 2,
    description: "Let the bot open while you play green.",
  },
];

interface HomeScreenProps {
  player: Player;
  onPlayerChange: (player: Player) => void;
  onGameStart: () => void;
}

function HomeScreen({ player, onPlayerChange, onGameStart }: HomeScreenProps) {
  return (
    <section className="home-screen" aria-labelledby="home-title">
      <div className="home-screen__intro">
        <span className="status-badge status-badge-info">New match</span>
        <h1 id="home-title">Choose your chip</h1>
        <p className="body-lg">
          Pick the color you want to play as before starting a calm match
          against the bot.
        </p>
      </div>

      <div className="home-screen__panel card">
        <div className="home-screen__preview" aria-hidden="true">
          <div className="home-screen__board">
            {Array.from({ length: 16 }).map((_, index) => (
              <span
                key={index}
                className={
                  index === 5 || index === 10
                    ? "home-screen__slot home-screen__slot--red"
                    : index === 6 || index === 9
                      ? "home-screen__slot home-screen__slot--yellow"
                      : "home-screen__slot"
                }
              />
            ))}
          </div>
        </div>

        <div className="home-screen__controls">
          <div className="home-screen__label">
            <h2>Select player color</h2>
            <p className="body-md">
              You can change this before the match begins.
            </p>
          </div>

          <div
            className="home-screen__options"
            role="radiogroup"
            aria-label="Player color"
          >
            {colorOptions.map((option) => (
              <button
                key={option.color}
                className={`color-option color-option--${option.color}`}
                type="button"
                role="radio"
                aria-checked={player === option.value}
                onClick={() => onPlayerChange(option.value)}
              >
                <span className="color-option__chip" aria-hidden="true" />
                <span>
                  <strong>{option.label}</strong>
                  <small>{option.description}</small>
                </span>
              </button>
            ))}
          </div>

          <button
            className="btn btn-primary home-screen__start"
            type="button"
            onClick={onGameStart}
          >
            Start game
          </button>
        </div>
      </div>
    </section>
  );
}

export default HomeScreen;
