# struct

## golang语言中struct的初始化方式

```go
// 先定义结构体
type Rect struct {
	width  int
	height int
}

// 再初始化结构体
rect1 := new(Rect)
rect2 := &Rect{}
rect3 := &Rect{10, 20}
rect4 := &Rect{width:10, height:20}

// 定义 + 初始化同时进行
rect5 := &struct{width int, height int}{10, 20}
```