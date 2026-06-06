package main

import (
	"syscall/js"
)

func add(a, b int) int {
	return a + b
}

func addWrapper(this js.Value, args []js.Value) any {
	if len(args) < 2 {
		return "Error: Missing Arguments"
	}

	num1 := args[0].Int()
	num2 := args[1].Int()

	return add(num1, num2)
}

func main() {
	c := make(chan struct{})

	js.Global().Set("goAdd", js.FuncOf(addWrapper))

	<- c
}
