export type Cell = 0 | 1 | 2;
export type Board = Cell[][];
export type Player = "Red" | "Yellow";

export type BoardState = {
  board: Board;
  player: Cell;
  winner: Cell;
  isDraw: boolean;
}

export type GameResponse = {
  data: BoardState | null;
  message: string;
  status: "success" | "error";
}