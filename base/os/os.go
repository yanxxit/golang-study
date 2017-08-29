package main

import (
	"os/exec"
	"fmt"
)

func main() {
	cmdo := exec.Command("ping", "127.0.0.1", "my first commond comes from golang")
	if err := cmdo.Start(); err != nil {
		fmt.Printf("Error:The command No.0 can not be startup:%s\n", err)
		return
	}
}
