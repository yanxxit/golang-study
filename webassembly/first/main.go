package main

import (
	"syscall/js"
)

func calFib(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func fib(params []js.Value) {
	value := params[0].Int()
	value = calFib(value)
    // 当前Go和wasm交互，wasm没法直接获得函数的返回值，调用window.wamsCallback(value)或者直接window.output获取
    // window.wamsCallback为用户在Javascript中自定义的函数，也就是一个回调函数
	js.Global().Set("output", value)
	js.Global().Call("wamsCallback", value)
}
 
// 将Go里面的方法注入到window.fibNative里面
func registerCallbacks() {
	js.Global().Set("fibNative", js.NewCallback(fib))
}

func main() {
	registerCallbacks()
	select {}
}