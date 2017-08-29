package main

import (
	"fmt"
)

func init() {
	var slice1 []int = make([]int, 10)

	fmt.Println(slice1)
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	show()
	getSlice()
	appendParams()
}

func show() {
	s := [] int{1, 2, 3, 433, 66}
	t := s[1:2]
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(s[1])
	fmt.Println("-------------")
	fmt.Println(len(s)) //切片是可索引的，并且可以由 len() 方法获取长度。
	fmt.Println(cap(s)) //切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

	fmt.Println("show")

	showNil()
}

func showNil() {
	var numbers []int
	if (numbers == nil) {
		fmt.Printf("切片是空的")
	}
}

func getSlice() {
	num := []string{"admin", "data"}
	fmt.Println(num)
	fmt.Println(num[0])
}

func appendParams() {
	fmt.Println("-----------------------------")
	var numbers []int
	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	fmt.Println(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)

	fmt.Println(numbers1)
}
