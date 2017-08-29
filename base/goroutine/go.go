package main

import (
	"fmt"
	//"runtime"
)

// go fmt go.go

func say(s string) {
	//for i := 0; i < 5; i++ {
	//	runtime.GOMAXPROCS(3) //使用多核，多线程
	//	runtime.Gosched()     //runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
	//	fmt.Println(s)
	//}
	fmt.Println(s)
}

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func getName(a string, c chan string) {
	name := a + "是一个有趣的人"
	c <- name

}

func main() {
	go say("world")
	say("hello")

	//使用channel
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)

	f := make(chan string)
	go getName("admin", f)
	go getName("root", f)
	m, n := <-f, <-f
	fmt.Println(m, n)

}
