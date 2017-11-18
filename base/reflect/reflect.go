package main

import (
	"fmt"
	"reflect"
)

//列举几个反射的例子：1）简单类型反射，2）复杂类型反射，3）对反射回来的数据的可修改属性

func demo1() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type :", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

type T struct {
	A int
	B string
}

// 2、复杂类型反射
func demo2() {
	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

//3、通过反射获得的变量的可设置属性
func demo3() {
	testA()
	testB()
}

func testA() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	if true == v.CanSet() {
		fmt.Println("v is setable")
		//v.Set(4.1)
	} else {
		fmt.Println("v is not setable")
	}
}

func testB() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settablitty of p :", p.CanSet())

	v := p.Elem()
	fmt.Println("settablitty of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

func main() {
	demo1()
	demo2()
	demo3()
}
