# Connect Four

A Connect Four game with an AI opponent. The bot uses a minimax algorithm with alpha-beta pruning. The frontend is built with React and TypeScript, while the AI logic runs as a Go WebAssembly module.

## Architecture

```
wasm/         Go AI engine compiled to WebAssembly
  main.go       JS-glue exports (startGame, makePlayerMove, etc.)
  game/
    board.go    Board types, dimensions, cell constants
    game.go     Game state, move handling, win detection
    bot.go      Minimax + alpha-beta pruning
    *_test.go   Unit tests

ui/           React + Vite + TypeScript frontend
  src/
    workers/go.worker.ts   Web Worker that loads and wraps the WASM module
    services/wasm_client.ts Comlink bridge to the worker
    components/
      Layout/       Fixed topbar with Help and Code icons
      HomeScreen/   Pre-game color selection
      GameScreen/   Board + status card (turn indicator / game over)
      HowToPlay/    Rules page with bento grid layout
```

## Build & Run

```bash
# 1. Compile the Go WASM module
cd wasm && GOOS=js GOARCH=wasm go build -o ../ui/public/main.wasm

# 2. Start the dev server
cd ../ui && npm run dev
```

## Test

```bash
go test ./game/...
```

## Deploy

The project deploys to GitHub Pages via a GitHub Actions workflow. On push to `main`, the workflow:

1. Compiles the Go WASM module
2. Copies the Go WASM bridge (`wasm_exec.js`)
3. Builds the React frontend with `--base=/connect-four/`
4. Uploads the output to GitHub Pages
