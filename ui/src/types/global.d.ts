import type { GameResponse } from "./game";

export { };

declare global {
  var Game: {
    startGame: () => GameResponse;
    makePlayerMove: (col: number) => GameResponse;
    makeBotMove: () => Promise<GameResponse>;
    resetGame: () => GameResponse;
  }
}