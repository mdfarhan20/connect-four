package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/mdfarhan20/connect-four/wasm/game"
)

type WasmResponse struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type GameState struct {
	Board  game.Board `json:"board"`
	Player game.Cell  `json:"player"`
	Winner game.Cell  `json:"winner"`
	IsDraw bool       `json:"isDraw"`
}

func parseJsonResponse(data interface{}) any {
	bytes, _ := json.Marshal(data)
	return js.Global().Get("JSON").Call("parse", string(bytes))
}

func jsonError(message string) any {
	err := WasmResponse{
		Data:    nil,
		Message: message,
		Status:  "error",
	}

	return parseJsonResponse(err)
}

var currentGame *game.Game

func startNewGame(this js.Value, args []js.Value) any {
	if currentGame != nil {
		return jsonError("A game is already in progress")
	}

	currentGame = game.NewGame()
	resp := WasmResponse{
		Data: GameState{
			Board:  currentGame.GetBoard(),
			Player: currentGame.GetPlayer(),
			Winner: 0,
			IsDraw: false,
		},
		Message: "Game started successfully",
		Status:  "success",
	}

	return parseJsonResponse(resp)
}

func makePlayerMove(this js.Value, args []js.Value) any {
	if currentGame == nil {
		return jsonError("There is no game in progress")
	}

	if len(args) < 1 {
		return jsonError("Column number is required to make a move")
	}

	if args[0].Type() != js.TypeNumber {
		return jsonError("Column number must be of type number")
	}

	column := args[0].Int()
	if column < 0 || column >= game.Width {
		return jsonError("Column must be within 0 and 6")
	}

	err := currentGame.MakeMove(column)
	if err != nil {
		return jsonError(err.Error())
	}

	resp := WasmResponse{
		Data: GameState{
			Board:  currentGame.GetBoard(),
			Player: currentGame.GetPlayer(),
			Winner: currentGame.Winner(),
			IsDraw: currentGame.IsDraw(),
		},
		Message: "Move made successfully",
		Status:  "success",
	}

	return parseJsonResponse(resp)
}

func makeBotMove(this js.Value, args []js.Value) any {
	promise := js.Global().Get("Promise")

	return promise.New(js.FuncOf(func(_ js.Value, pArgs []js.Value) any {
		resolve := pArgs[0]
		reject := pArgs[1]

		go func() {
			defer func() {
				if r := recover(); r != nil {
					reject.Invoke(jsonError(fmt.Sprintf("Go Runtime Panic: %v", r)))
				}
			}()

			if currentGame == nil {
				reject.Invoke(jsonError("There is no game in progress"))
				return
			}

			if currentGame.IsDraw() || currentGame.Winner() != game.Empty {
				reject.Invoke(jsonError("This game has ended"))
				return
			}

			move := game.FindBestMove(*currentGame)
			currentGame.MakeMove(move)

			resp := WasmResponse{
				Data: GameState{
					Board:  currentGame.GetBoard(),
					Player: currentGame.GetPlayer(),
					Winner: currentGame.Winner(),
					IsDraw: currentGame.IsDraw(),
				},
				Message: "Bot move made successfully",
				Status:  "success",
			}

			resolve.Invoke(parseJsonResponse(resp))
		}()

		return nil
	}))
}

func resetGame(this js.Value, args []js.Value) any {
	if currentGame == nil {
		return jsonError("There is no game in progress")
	}

	currentGame = game.NewGame()
	resp := WasmResponse{
		Data: GameState{
			Board:  currentGame.GetBoard(),
			Player: currentGame.GetPlayer(),
			Winner: 0,
			IsDraw: false,
		},
		Message: "Game reset successfully",
		Status:  "success",
	}

	return parseJsonResponse(resp)
}

func main() {
	c := make(chan struct{})

	gameApi := js.ValueOf(map[string]interface{}{
		"startGame":      js.FuncOf(startNewGame),
		"makePlayerMove": js.FuncOf(makePlayerMove),
		"makeBotMove":    js.FuncOf(makeBotMove),
		"resetGame":      js.FuncOf(resetGame),
	})

	js.Global().Set("Game", gameApi)

	<-c
}
