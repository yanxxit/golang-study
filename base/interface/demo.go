package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func TransInterfaceToString(e interface{}) string {
	var str string
	switch v := e.(type) {
	case int:
		fmt.Println("整型", v)
		var s int
		s = v
		fmt.Println(s)
	case string:
		fmt.Println("字符串", v)
		str = v
	}
	return str
}

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}
func (s *S) Set(age int) {
	s.Age = age
}

//3
func f(i I) {
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	//测试功能
	s := S{}
	f(&s)  //4

}
