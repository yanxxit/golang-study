package timeutil

import (
	"testing"
	"fmt"
	"time"
)

func TestRandomString(t *testing.T) {
	fmt.Println(time.Now().Unix())
	for i := 0; i < 10; i++ {
		start := time.Now()
		fmt.Println("生成随机数：" + RandomString(3))
		end := time.Now()
		delta := end.Sub(start)
		fmt.Printf("程序平均执行时间：%s\n", delta)
	}
}
