package main

import (
	"strconv"
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {

	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	println(int1 + int2)
	// js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
	return int1 + int2

}

func subtract(this js.Value, i []js.Value) interface{} {

	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	println(int1 - int2)
	// js.Global().Set("output", js.ValueOf(i[0].Int()-i[1].Int()))
	return js.ValueOf(int1 - int2).String()

}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)
	println("WebAssembly Initialized")
	registerCallbacks()

	<-c
}
