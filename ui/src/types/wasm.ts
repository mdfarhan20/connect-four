import type { GameResponse } from "./game"

export interface WasmApi {
  ready(): Promise<boolean>;
  startGame(): Promise<GameResponse>;
  makePlayerMove(col: number): Promise<GameResponse>;
  makeBotMove(): Promise<GameResponse>;
  resetGame(): Promise<GameResponse>;
}