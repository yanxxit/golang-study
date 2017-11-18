package main

import (
	"fmt"
	"os"
	"os/exec"
)

func init() {
	fmt.Println("pid", os.Getpid())
	fmt.Println("ppid", os.Getppid())
}

func main() {
	cmdo := exec.Command("echo", "-n", "my first commond comes from golang")
	fmt.Println(cmdo)
	if err := cmdo.Start(); err != nil {
		fmt.Printf("Error:The command No.0 can not be startup:%s\n", err)
		return
	}

	stdout0, err := cmdo.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:The command No.0 can not be startup:%s\n", err)
		return
	}
	output0 := make([]byte, 30)
	n, err := stdout0.Read(output0)
	fmt.Printf("%s\n", output0[0:n])

}
