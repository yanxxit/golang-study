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

func main() {
	//c := make(chan int)
	//go send(c)
	//go recv(c)
	//time.Sleep(3 * time.Second)
	//close(c)
	c := make(chan int, 10)
	c1 := make(chan string, 100)
	// go send(c)
	// go recv(c)
	go sendS(c1)
	go recvS(c1)
	time.Sleep(3 * time.Second)
	close(c)
}
