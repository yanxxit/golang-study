/* 删除文件函数
   os.Remove(file string) error

   file  文件名
   error 如果失败则返回错误信息

*/

package main

import (
	"fmt"
	"os"
)

func main() {
	file := "test.txt"     //源文件路径
	err := os.Remove(file) //删除文件test.txt
	if err != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
		//如果删除成功则输出 file remove OK!
		fmt.Print("file remove OK!")
	}

}
