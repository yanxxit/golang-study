# webassembly
https://www.jianshu.com/p/540ab3db556e

GOARCH=wasm GOOS=js go build -o test.wasm main.go

GOARCH=wasm GOOS=js go build -o main.wasm main.go

## add
```go
// function definition
func add(this js.Value, i []js.Value) interface{} {
  return js.ValueOf(i[0].Int()+i[1].Int())
}

// exposing to JS
js.Global().Set("add", js.FuncOf(add))
```

## 参考
- https://asyncoder.com/2018/08/27/%E5%88%A9%E7%94%A8Golang%E7%8E%A9%E8%BD%ACWebAssembly/
- https://zhuanlan.zhihu.com/p/43506954
- 利用Golang玩转WebAssembly http://asyncoder.com/2018/08/27/%E5%88%A9%E7%94%A8Golang%E7%8E%A9%E8%BD%ACWebAssembly/