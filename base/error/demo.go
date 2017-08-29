package main

import (
	"github.com/henrylee2cn/faygo/errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf:elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}
