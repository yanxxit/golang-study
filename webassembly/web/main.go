// main.go
package main
import (
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
  return js.ValueOf(i[0].Int()+i[1].Int())
}


func changeBodyColor(color []js.Value) {
	// document.body.style.color = color
	js.Global().Get("document").Get("body").Set("style", "color:"+color[0].String())
}

func setInputValue(val []js.Value) {
	id := val[0].String()
	// document.getElementById(id).value = "value from Go"
	js.Global().Get("document").Call("getElementById", id).Set("value", "value from Go")
}

func registerCallbacks() {
	// js.Global().Set("sum", js.NewCallback(sum))
	js.Global().Set("sum", js.FuncOf(add))
	js.Global().Set("changeBodyColor", js.NewCallback(changeBodyColor))
	js.Global().Set("setInputValue", js.NewCallback(setInputValue))
}

func main() {
	c := make(chan struct{}, 0)
	println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}
