package main

import "fmt"

type MyInterface interface {
	Print()
}

func TestFunc(x MyInterface) {

}

type MyStruct struct {
}

func (p MyStruct) Print() {

}

//声明一个USB的接口
type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (p PhoneConnecter) Name() string {
	return p.name
}

func (p PhoneConnecter) Connect() {
	fmt.Println(p.Name(), "已连接成果")
}

type PcConnecter struct {
	title string
}

func (p PcConnecter) Name() string {
	return p.title
}

func (p PcConnecter) Connect() {
	fmt.Println(p.Name(), "已连接成果")
}

func main() {
	var me MyStruct
	TestFunc(me) //只要拥有相同的方法，就算一个就可以被一个接口使用

	var u USB                   //声明u的类型是USB类型，interface
	var u1 USB                  //声明u的类型是USB类型，interface
	u = PhoneConnecter{"小明的手机"} //可以直接给符合条件的interface 赋值
	u1 = PcConnecter{"小明的电脑"}   //可以直接给符合条件的interface 赋值
	u.Connect()
	u1.Connect()
}
