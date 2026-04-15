# Connect Four Bot

## Overview
Connect Four game with AI opponent using minimax algorithm.

## Directory Structure
- `wasm/` - Go WASM bot module
  - `main.go` - WASM entry point
  - `game/` - Core game logic
    - `board.go` - Board types, dimensions, cell constants
    - `game.go` - Game state, move handling, win detection
    - `bot.go` - Minimax algorithm for AI moves
    - `utils.go` - Helper functions
- `src/` - React + Vite + TypeScript UI (planned)

## Go Module
- Module: `github.com/mdfarhan20/connect-four/wasm`
- Packages: `game`

## Build Commands
- `cd wasm && go build -o game.wasm` (requires WASM target)
- TinyGo recommended: `tinygo build -o game.wasm -target wasm ./...`

## Bot Architecture
- **Algorithm**: Minimax (no alpha-beta pruning yet)
- **Board**: 6 rows x 7 columns
- **Cells**: `Empty(0)`, `Red(1)`, `Yellow(2)`
- **Function**: `FindBestMove(g Game) int`

## Future
- React + Vite + TypeScript UI
- WASM integration with frontend
