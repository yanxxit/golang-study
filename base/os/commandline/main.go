/**
用os包和flag包实现读取main命令入参的相关资料
os包的Args包括了命令名本身，作为读取参数的第一个值。
flag.Args读取命令后的入参（要配合flag.Parse()使用）。
也就是，os的比flag多一个命令名的读取。
from:http://www.jb51.net/article/63015.htm
 */
package main

import (
	"os"
	"fmt"
	"flag"
	"bufio"
)

func main() {
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	fmt.Println(line, hasMoreInLine, err)
	// os.Args方式
	args := os.Args
	if args == nil || len(args) < 2 {
		fmt.Println("Hello 世界!")
	} else {
		fmt.Println("Hello ", args[1]) // 第二个参数，第一个参数为命令名
	}

	// flag.Args方式
	flag.Parse()
	var ch []string = flag.Args()
	if ch != nil && len(ch) > 0 {
		fmt.Println("Hello ", ch[0]) // 第一个参数开始
	}
}
