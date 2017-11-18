package main

//需要验证
import (
	"C"
	"fmt"
)

/*
#include <stdlib.h>
#include <stdio.h>
void hello()
{
	printf("Hello World 1111 !\n");
}
*/

func Hello() {
	C.hello()
}

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(100)
	fmt.Println("Random:", Random())
	Hello()
}
