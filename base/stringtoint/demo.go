package main

import (
	"strconv"
	"fmt"
)

/**
#string到int
int,err:=strconv.Atoi(string)
#string到int64
int64, err := strconv.ParseInt(string, 10, 64)
#int到string
string:=strconv.Itoa(int)
#int64到string
string:=strconv.FormatInt(int64,10)
 */

func main()  {
	var num int
	num ,_= strconv.Atoi("10")
	fmt.Println(num)
}
