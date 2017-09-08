package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	fmt.Println(os.Args)
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入您的名字：")

	//读取数据知道
	input, err := inputReader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("hello,%s what can I do for you ?", name)
		//os.Stdout(name);
	}
	for true {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		} else {
			name := input[:len(input)-1]
			fmt.Printf("hello,%s ! what can I do for you ?", name)

			if name == "exit" {
				os.Exit(1)
			}
			//os.Stdout(name);
		}
	}

}
