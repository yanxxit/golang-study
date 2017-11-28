// +----------------------------------------------------------------------
// | 文件标题
// | Created by InteliJ IDE
// +----------------------------------------------------------------------
// | 别忘了描述
// +----------------------------------------------------------------------
// | Copyright (c) 2017 http://www.mohoo.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | @Author: yxxit <yanxx@infoepoch.com>
// | @Date  : 2017/11/22
// | @Time  : 8:51
// +----------------------------------------------------------------------
package main

import "fmt"

func main() {
	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic(55)
	fmt.Println("b")
	fmt.Println("f")
}
