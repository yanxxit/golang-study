package main

import (
	"fmt"
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

//只能接收channel中的数据
func recv(c <-chan int) {
	for i := range c {
		fmt.Println("received", i)
	}
}

func main() {
	//c := make(chan int)
	//go send(c)
	//go recv(c)
	//time.Sleep(3 * time.Second)
	//close(c)
	c := make(chan int,10)
	go send(c)
	go recv(c)
	time.Sleep(3 * time.Second)
	close(c)
}
