// B"H

package main

import (  
	"syscall/js"
	"strconv"
)

// func add(i []js.Value) {
//     js.Global().Set( "output", js.ValueOf( i[0].Int() + i[1].Int() ) )
//     println(js.ValueOf(i[0].Int() + i[1].Int()).String())
// }

// function definition
func add(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(i[0].Int()+i[1].Int())
}

// function definition
func subtract(this js.Value, i []js.Value) interface{} {
	return js.ValueOf(i[0].Int()-i[1].Int())
}


func addInDom(this js.Value, i []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1 + int2)
	return int1 + int2
}


func subtractInDom(this js.Value, i []js.Value) interface{} {
    value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
    value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

    int1, _ := strconv.Atoi(value1)
    int2, _ := strconv.Atoi(value2)

    js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1 - int2)
	return int1 - int2
}


func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))
    js.Global().Set("subtract", js.FuncOf(subtract))
	js.Global().Set("addInDom", js.FuncOf(addInDom))
	js.Global().Set("subtractInDom", js.FuncOf(subtractInDom))
}

func main() {
    c := make(chan struct{}, 0)

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    <-c
}