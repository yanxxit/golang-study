package main

import "fmt"

type One interface {
	Say() string
}

type Two interface {
	Say2() string
}

/**
多个接口，合并成一个接口
 */
type More interface {
	One
	Two
	GetName() string
}
// person
type Person struct {
	Name string
	Age  int32
}

func (p *Person) Say() string {
	fmt.Println(p.Name + "说：hello world ")
	return "hello world"
}
func (p *Person) Say2() string {
	return "hello world2"
}
func (p *Person) GetName() string {
	fmt.Println(p.Name)
	return p.Name
}
func main() {
	var a More
	a = &Person{Name: "张三", Age: 123}
	a.GetName()
	a.Say()
}
