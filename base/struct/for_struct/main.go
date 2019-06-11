package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

/**
遍历结构体字段：
1. 通过reflect 进行处理
*/
func main() {
	v := reflect.ValueOf(person{"steve", 30})
	count := v.NumField()
	for i := 0; i < count; i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			fmt.Println(f.String())

		case reflect.Int:
			fmt.Println(f.Int())
		}
	}
}
