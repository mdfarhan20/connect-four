export type Cell = 0 | 1 | 2;
export type Board = Cell[][];
export type Player = 1 | 2;

export type BoardState = {
  board: Board;
  player: Player;
  winner: Cell;
  isDraw: boolean;
};

export type GameResponse = {
  data: BoardState | null;
  message: string;
  status: "success" | "error";
};
