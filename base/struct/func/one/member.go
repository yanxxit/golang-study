package main

import "fmt"

type Member struct {
	Id     int    `json:"id,-"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender int    `json:"gender,"`
	Age    int    `json:"age"`
}

func setName(m Member, name string) { //普通函数
	m.Name = name
}

/**
结构体是值传递，当我们调用setName时，方法接收器接收到是只是结构体变量的一个副本，通过副本对值进行修复，并不会影响调用者，因此，我们可以将方法接收器定义为指针变量，就可达到修改结构体的目的了。
*/
func (m Member) setName(name string) { //绑定到Member结构体的方法
	m.Name = name
}

/**
绑定到Member结构体的方法
*/
func (m *Member) setEmail(email string) {
	m.Email = email
}

func MemberRun() {
	m := Member{}
	m.setName("小米")
	m.setEmail("123@qq.com")
	fmt.Println("----->", m.Name)
	fmt.Println("*****>", m.Email)
}

type Animal struct {
	Name   string  //名称
	Color  string  //颜色
	Height float32 //身高
	Weight float32 //体重
	Age    int     //年龄
}

//奔跑
func (a Animal) Run() {
	fmt.Println(a.Name + "is running")
}

//吃东西
func (a Animal) Eat() {
	fmt.Println(a.Name + "is eating")
}

type Lion struct {
	Animal Animal //匿名字段
	//Animal //匿名字段
	About string
}

type Cat struct {
	Animal
}

func Golion() {
	var lion = Lion{
		Animal: Animal{Name: "小狮子", Color: "灰色"}, About: "aaaa",
	}
	var cat = Cat{
		Animal{Name: "kitty", Color: "黑色"},
	}
	//匿名:猫
	cat.Eat()
	cat.Run()
	//非匿名：狮子
	lion.Animal.Run()
	fmt.Println(lion.About)
}

func main() {
	MemberRun()
	Golion()
}
