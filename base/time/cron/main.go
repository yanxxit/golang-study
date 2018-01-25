// +----------------------------------------------------------------------
// | 文件标题
// | Created by InteliJ IDE
// +----------------------------------------------------------------------
// | 别忘了描述
// +----------------------------------------------------------------------
// | Copyright (c) 2017 http://www.mohoo.com.cn All rights reserved.
// +----------------------------------------------------------------------
// | @Author: yxxit <yanxx@infoepoch.com>
// | @Date  : 2017/12/15
// | @Time  : 14:07
// +----------------------------------------------------------------------
package main

import (
	"time"
	"fmt"
)

func main()  {
	ticker := time.NewTicker(time.Minute * 1)
	go func() {
		for _ = range ticker.C {
			fmt.Printf("ticked at %v", time.Now())
		}
	}()
}
