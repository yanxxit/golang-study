package main

import (
	"os/exec"
	"bytes"
	"fmt"
)

func exec_shell() {
	cmd := exec.Command("git", "pull")
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("--",err)
	}
	fmt.Printf("abc:-%s", out.String())
}
func main() {
	exec_shell()
}
