package main

import (
	"fmt"
	"sync"
	"time"
)

// StartAsync WaitGroup简称同步组，用于等待goroutines结束的。
// StartAsync demo
func StartAsync() {

}

func main() {

	t1 := time.Now() // get current time
	var wg sync.WaitGroup

	wg.Add(2) // 因为有两个动作，所以增加2个计数
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1")
		wg.Done() // 操作完成，减少一个计数
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Goroutine 2")
		wg.Done() // 操作完成，减少一个计数
	}()

	wg.Wait() // 等待，直到计数为0
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)

	t2 := time.Now() // get current time
	var wg2 sync.WaitGroup

	wg2.Add(2) // 因为有两个动作，所以增加2个计数
	list := []string{"349639", "100042"}
	for i, userid := range list {
		fmt.Println(i, userid)
		go func(userid string) {
			time.Sleep(1 * time.Second)
			fmt.Println("Goroutine 1", userid)
			wg2.Done() // 操作完成，减少一个计数
		}(userid)
	}
	wg2.Wait() // 等待，直到计数为0
	elapsed1 := time.Since(t2)
	fmt.Println("App elapsed: ", elapsed1)
}
