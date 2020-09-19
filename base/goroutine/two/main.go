package main

import (
	"fmt"
	"strconv"
	"time"
)

//只能向chan里send数据
func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send readey ", i)
		c <- i
		fmt.Println("send ", i)
	}
}

func sendS(c chan<- string) {
	for i := 0; i < 100; i++ {
		tt := strconv.Itoa(i)
		fmt.Println("- send readey ", tt)
		c <- tt + ":---"
		fmt.Println("- send ", tt)
	}
}

//只能接收channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received", i)
	}
}

func recvS(c <-chan string) {
	for i := range c {
		fmt.Println("received", i)
	}
}
func test(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ", i)
		c <- i
	}
}
func test1(c chan interface{}) {
	for i := 0; i < 10; i++ {
		fmt.Println("send ", i)
		resData := make(map[string]interface{})
		resData["admin"] = "yanxxit"
		resData["test"] = "test"
		resData["pwd"] = "123456"
		resData["age"] = 12
		resData["index"] = i
		c <- resData
	}
}
func main() {
	//c := make(chan int)
	//go send(c)
	//go recv(c)
	//time.Sleep(3 * time.Second)
	//close(c)
	c := make(chan int, 10)
	c1 := make(chan string, 100)
	ch := make(chan int)
	ch1 := make(chan interface{})
	go test(ch)
	go test1(ch1)
	for j := 0; j < 10; j++ {
		fmt.Println("get ", <-ch)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("get---- ", <-ch1)
	}
	// go send(c)
	// go recv(c)
	// go sendS(c1)
	// go recvS(c1)
	close(c)
	close(c1)
	close(ch)
	close(ch1)
	time.Sleep(3 * time.Second)

	list := []string{"349639", "100042"}
	for i, userid := range list {
		fmt.Println(i, userid)

	}
}
