package main

import (
	"fmt"
	"strings"
)

//string 替换功能
func main() {
	str := "welcome to sharejs.com"
	info := "welcom<br/>to<br/>sharejs.com"
	str = strings.Replace(str, " ", ",", -1)
	info = strings.Replace(info, "<br/>", "- ", -1)
	fmt.Println(str)
	fmt.Println(info)
}
