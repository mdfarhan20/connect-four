# Connect Four Bot

## Overview
Connect Four game with AI opponent using minimax algorithm, integrated with a React frontend using Go compiled to WebAssembly (WASM).

## Directory Structure
- `wasm/` - Go WASM bot module
  - `go.mod` - Go module file (Go 1.23.3)
  - `main.go` - WASM entry point (exposes the global `Game` object to JS)
  - `game/` - Core game logic
    - `board.go` - Board types, dimensions, cell constants, `Point` and `Direction` helper types
    - `board_test.go` - Tests for board/point validation and movement
    - `game.go` - Game state, move handling, win detection
    - `game_test.go` - Tests for game functions (moves, wins, draws)
    - `bot.go` - Minimax algorithm with alpha-beta pruning for AI moves
    - `bot_test.go` - Tests for bot move calculation and blocking logic
    - `utils.go` - Helper functions (board printer, board initialization)
    - `utils_test.go` - Tests for utility functions
- `ui/` - React + Vite + TypeScript UI frontend
  - `package.json` - Node dependencies (React 19, Vite 8, Comlink) and scripts (`dev`, `build`, `lint`, `preview`)
  - `index.html` - Application HTML entry point
  - `tsconfig.json` / `tsconfig.app.json` / `tsconfig.node.json` - TypeScript configuration
  - `vite.config.ts` - Vite configurations
  - `eslint.config.js` - ESLint configuration
  - `public/` - Static assets directory
    - `main.wasm` - Compiled Go WebAssembly binary
    - `wasm_exec.js` - Go standard WASM bridge runtime
  - `src/` - React application source code
    - `main.tsx` - React mount entry point
    - `App.tsx` - Main React component and screen composition after WASM loading
    - `index.css` - Global stylesheet
    - `components/` - Reusable UI components
      - `Loader/` - WASM loading state component with colocated CSS
      - `Layout/` - Shared app shell and top bar used across screens
      - `HomeScreen/` - Pre-game color selection screen
    - `constants/` - Game constants (`game.ts`)
    - `types/` - TypeScript type declarations
      - `game.ts` - Game data types (Cell, Board, Player)
      - `global.d.ts` - Typing for global WASM methods under `globalThis.Game`
      - `wasm.ts` - Comlink-wrapped Web Worker API typings
      - `wasm_exec.d.ts` - Declarations for the Go WASM runtime class
    - `scripts/` - Script assets (holds a copy of `wasm_exec.js`)
    - `services/` - Services
      - `wasm_client.ts` - Initializes and exposes the Comlink-wrapped WASM Web Worker client
    - `workers/` - Web Workers
      - `go.worker.ts` - Asynchronous worker running the Go WASM instance

## Go Module
- Module: `github.com/mdfarhan20/connect-four/wasm`
- Packages: `main`, `game`

## Build Commands
- **Compile Go WASM module**: 
  - Windows PowerShell:
    ```powershell
    $env:GOOS="js"; $env:GOARCH="wasm"; go build -o ../ui/public/main.wasm
    ```
  - Linux/macOS:
    ```bash
    GOOS=js GOARCH=wasm go build -o ../ui/public/main.wasm
    ```
- **TinyGo recommended compile**: 
  - `tinygo build -o ../ui/public/main.wasm -target wasm ./...`
- **UI Development**: 
  - `npm run dev` (run from the `ui/` directory)
- **UI Production Build**: 
  - `npm run build` (run from the `ui/` directory)

## Test Commands
- `go test ./game/...` - Run core game logic tests (executable on host machine)
- `go test -v ./game/...` - Verbose output for game package tests
- *Note*: Running `go test ./...` directly on a non-WASM host fails because the `main` package imports `syscall/js`. To test everything including `main.go`, you must run under a WASM environment (e.g. node with wasm runner).

## Bot Architecture
- **Algorithm**: Minimax with alpha-beta pruning bounds check
- **Board**: 6 rows x 7 columns
- **Cells**: `Empty(0)`, `Red(1)`, `Yellow(2)`
- **Exposed Go Functions**: `FindBestMove(g Game) int`
- **WASM Global JS Functions** (exposed via `globalThis.Game`):
  - `startGame()` - Starts a new game instance
  - `makePlayerMove(col)` - Makes a player move in the specified column
  - `makeBotMove()` - Computes and makes a bot move (returns Promise)
  - `resetGame()` - Resets the current game

## UI Design System & Theme
- **Design System**: Soft Play (`asset-stub-assets-4832b5f3762c414ba41873574216f659-1777211386496`)
- **Aesthetic**: Soft minimalism with pillowy/tactile cues, desaturated tones to minimize eye strain.
- **Color Palette**:
  - Background & Surfaces: `#fff8f1` (warm off-white/cream) with a radial dot pattern grid (`24px` interval)
  - Primary (Player 1 / Red): `#8c4e33` (sun-baked terracotta)
  - Secondary (Player 2 / Yellow): `#416465` (dusty sage green)
  - Tertiary (Accent/Meta): `#0d6969` (supportive deep teal)
- **Typography**: **Plus Jakarta Sans** imported via Google Fonts.
- **Elevation & Shapes**: Super-ellipses (`0.5rem` to `1.5rem` border-radius) for layout cards/buttons. Perfect circles for chips/board holes. Uses ambient/tinted diffused glows and board hole inset shadows.
- **Interactions**: Tactile button presses (bottom-border offsets translating `2px` down on click) and smooth chip-drop bounce animations.

## Stitch UI Workflow
- When a Stitch URL is provided for a screen, use the Stitch MCP tools to fetch the project and target screen before implementing the UI.
- Parse the project ID from `/projects/{projectId}` and the screen instance ID from `node-id={screenInstanceId}`.
- Use `get_project` to find the matching `screenInstances` entry, then call `get_screen` with that entry's `sourceScreen`.
- Recreate the screen as React UI only. Do not wire gameplay behavior unless explicitly requested.
- Keep reusable chrome such as the top bar in layout components, and colocate screen/component CSS with the component instead of growing `index.css`.
- Preserve the WASM integration flow: show the loader while `initializeWasm()` resolves, then render the current UI screen.

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
