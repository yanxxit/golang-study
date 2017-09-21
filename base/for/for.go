package main

import (
	"fmt"
)

func forMap() {
	fmt.Println("循环Map")
	list := make(map[string]interface{})
	list["openid"] = "1111111111111111111"
	list["deviceno"] = "2222222222222222222"
	list["bpnum"] = "333333333333333333333333333"
	for k, v := range list {
		fmt.Println(k, v)
	}
}

func forSlice() {
	s := []int{}
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	s = append(s, 10)

	for k, v := range s {
		fmt.Println(k, v)
	}

}

func main() {
	forMap()
	forSlice()
}
