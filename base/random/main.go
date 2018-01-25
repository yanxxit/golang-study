package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		fmt.Println(r.Intn(100))
	}
	fmt.Println("-------------------------")
	fmt.Println(r.Intn(100))
}
