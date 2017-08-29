package main

import (
	"fmt"
)

type User struct {
	Name string
	Age  int64
	Pwd  string
}

func (u *User) getBase() {
	fmt.Println(u.Name, u.Age, u.Pwd)
}

func main() {
	var user User

	user.Name = "张三"
	user.Age = 12
	user.Pwd = "12312312"

	user.getBase()

	name := getName("15806111230")
	res := getName("158061112301")
	num := add(3)
	fmt.Println("打印结果：" + name)
	fmt.Println("num：")
	fmt.Println("打印结果：" + res)
	fmt.Println("项目启动")

	fmt.Println(num)

	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}

/**
对设备号加密
 */
func getName(deviceno string) string {
	if len(deviceno) == 11 {
		return deviceno
	} else {
		return "不符合条件"
	}
}

func add(num int) int {
	return num * 10
}

func swap(x, y string) (string, string) {
	return y, x
}
