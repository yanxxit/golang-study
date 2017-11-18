package main

import (
	"fmt"
	"os"
)
//这段代码是试图打开指定的文件，然后判断是否有错误
//Go语言判断指定的文件是否存在
func main() {
	f, err := os.Open("dotcoo.com.txt")
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("file not exist!\n")
		return
	}
	fmt.Printf("file exist!\n")
	defer f.Close()
}
//该代码片段来自于: http://www.sharejs.com/codes/go/4367