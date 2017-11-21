package main

import "fmt"

type User struct {
	Name   string
	Pwd    string
	Age    int
	Love   string
	Mobile string
	Email  string
	Openid string
}

//获取用户名称
func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetNickName(openid string) string {
	return openid + "的博客"
}

func main() {
	my := &User{Name: "admin", Pwd: "123123", Mobile: "17721021494", Email: "123@qq.com", Openid: "abcdefg"}
	name := my.GetName()
	openid := my.GetNickName("openid...")
	fmt.Println(name)
	fmt.Println(openid)
	testDerive()
}

//1.golang的继承是继承父结构体的所有属性和方法，包括大小写开头的变量和函数。
//2.如果子结构体和父结构体有同名的变量或者函数，并不会产生覆盖，可以通过“父结构名字.变量或函数名”的方式调用父结构体的同名变量或函数
type A struct {
	aa int
	BB string
}
type C struct {
	Name   string
	mobile string
}

type B struct {
	A
	C
	aa int
	CC string
}

func (a *A) aFunc() {
	fmt.Println("A.aFunc")
}

func (a *A) BFunc() {
	fmt.Println("A.BFunc")
}

func (b *B) aFunc() {
	fmt.Println("B.aFunc")
}

func (b *B) CFunc() {
	fmt.Println("B.CFunc")
}

func testDerive() {
	var tb B
	//变量测试
	tb.aa = 9
	tb.A.aa = 8
	//如果子结构体 和父结构体有相同的参数，互补干扰
	fmt.Println("tb.aa", tb.aa)
	fmt.Println("tb.A.aa", tb.A.aa)

	tb.BB = "A.BB"
	tb.CC = "B.CC"
	tb.CC = "B.CC"
	tb.Name = "张山"
	tb.C.Name = "李四"
	tb.C.mobile = "17721021494"
	tb.mobile = "1580611230"
	//如果子结构体和父结构体没有相同的参数 则B.C.Name ==B.Name
	fmt.Println("tb.Name", tb.Name)
	fmt.Println("tb.C.Name", tb.C.Name)
	fmt.Println("tb.mobile", tb.mobile)
	fmt.Println("tb.C.mobile", tb.C.mobile)

	fmt.Println(tb.aa)
	fmt.Println(tb.A.aa)
	fmt.Println(tb.BB)
	fmt.Println(tb.A.BB)
	fmt.Println(tb.CC)

	//函数测试
	//tb.aFunc()
	//tb.A.aFunc()
	//tb.BFunc()
	//tb.A.BFunc()
	//tb.CFunc()
}
