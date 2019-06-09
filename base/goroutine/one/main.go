package main

import (
	"fmt"
	"golang-study/base/goroutine/one/wait"
	"runtime"
	"time"
)

func sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i * 2
	}
	time.Sleep(2 * time.Second)
	fmt.Println(sum)
}

func main() {
	//获取当前的GOMAXPROCS
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	//设置GOMAXPROCS的值为2
	runtime.GOMAXPROCS(2)
	//获取当前的GOMAXPROCS的值
	runtime.GOMAXPROCS(0)

	//引用外部
	wait.GetMore()

	fmt.Printf("hello %s", "world")
	//go sum()
	//匿名函数
	//go func() {
	//	sum := 0
	//	for i := 0; i < 10000; i++ {
	//		sum += i
	//	}
	//
	//	time.Sleep(1 * time.Second)
	//	fmt.Println((sum))
	//}()

	//NumGoruntine 可以返回当前程序的gorountine 数目
	fmt.Println("Num Go runtine=", runtime.NumGoroutine())

	//使用chan 进行进程间通讯
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)
		//写通道
		c <- struct{}{}
	}(c, ci)

	//读通道c,通过通道进行同步等待
	<-c

	//此时ci通道已经关闭，匿名函数启动的goruntine已经退出
	fmt.Println("NumGoroutine=", runtime.NumGoroutine())

	//但通道ci还可以继续读取
	for v := range ci {
		fmt.Println(v)
	}

	//main go runtine故意 sleep 5秒，防止其意外退出
	time.Sleep(5 * time.Second)

}
