package main

import (
	"fmt"
	"strings"
	"net/url"
)

func demo1() {
	s := "hello"
	c := []byte(s) //将字符串s转换为[]byte类型
	fmt.Println(c)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)
}

func demo2() {
	s := "hello"
	s = "demo_" + s[1:] //字符串虽然不能修改，但可以进行切片操作
	fmt.Printf("%s\n", s)
}

func demo3() {
	//` 括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出
	m := `hello
	world`
	fmt.Printf("%s\n", m)
}

func init() {
	demo1()
	demo2()
	demo3()
}

func main() {
	urldemo()
	strdemo()
}

func urldemo() {
	fmt.Println("--------------------------- url")
	base := "http%3A%2F%2Fm.sh.189.cn%2Fservice%2FbackPayInfo%3Fid%3D2016061297349034995730725"
	fmt.Println(base)
	str1, _ := url.QueryUnescape(base)
	fmt.Println(str1) //http://m.sh.189.cn/service/backPayInfo?id=2016061297349034995730725
	str2 := url.QueryEscape(str1)
	fmt.Println(str2) //http%3A%2F%2Fm.sh.189.cn%2Fservice%2FbackPayInfo%3Fid%3D2016061297349034995730725
}

func strdemo() {
	fmt.Println("------------------- strings 常用判断")
	str := "Hello world"
	fmt.Println("初始值：", str)
	fmt.Println("字符长度", len(str))

	fmt.Println("转换为小写：", strings.ToLower(str))
	fmt.Println("转换为大写：", strings.ToUpper(str))
	fmt.Println("比较：", strings.Compare(str, "Hel"))
	fmt.Println("比较：", strings.Compare(str, "world"))
	fmt.Println("比较：", strings.EqualFold(str, "world"))
	fmt.Println("比较：", strings.EqualFold(str, "Hello world"))

	fmt.Println("判断字符串是否包含“e”:", strings.Contains(str, "e"))
	prize := "https://m.sh.189.cn/telecoupon/getCoupon/page?couponid=4262C910-01F6-C0DF-A594-A1D9C61EB04F"

	fmt.Println(prize)
	fmt.Println("分割：", strings.Split(prize, "couponid="))
	fmt.Println("分割：", strings.Split(prize, "couponid=")[1])

}
