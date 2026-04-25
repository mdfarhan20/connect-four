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
    - `utils_test.go` - Tests for utils functions
    - `game_test.go` - Tests for game functions
- `src/` - React + Vite + TypeScript UI (planned)

## Go Module
- Module: `github.com/mdfarhan20/connect-four/wasm`
- Packages: `game`

## Build Commands
- `go build -o game.wasm` (requires WASM target)
- TinyGo recommended: `tinygo build -o game.wasm -target wasm ./...`

## Test Commands
- `go test ./...` - Run all tests
- `go test -v ./...` - Verbose output
- `go test -v -run TestName ./...` - Run specific test

## Bot Architecture
- **Algorithm**: Minimax (no alpha-beta pruning yet)
- **Board**: 6 rows x 7 columns
- **Cells**: `Empty(0)`, `Red(1)`, `Yellow(2)`
- **Function**: `FindBestMove(g Game) int`

## Testing Conventions
- **Test file location**: `wasm/game/` directory (same as source)
- **Test file naming**: `<filename>_test.go`
- **Package**: `package game` (to access unexported functions)
- **Test organization**: Use `t.Run()` subtests for multiple cases

## Git Commit Conventions
- **Format**: Lowercase with `action: description` structure (no commas)
- **Examples**:
  - `fix: correct game logic for move validation`
  - `test: add unit tests for board functions`
  - `docs: update API documentation`
- **Action Types**: Use `fix`, `feat`, `docs`, `test`, `refactor`, `chore`
- **Best Practice**: Separate concerns into focused commits when possible

## Future
- React + Vite + TypeScript UI
- WASM integration with frontend
