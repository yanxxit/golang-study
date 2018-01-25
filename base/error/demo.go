package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	err := errors.New("emit macho dwarf:elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err.Error())
	fmt.Println(reflect.TypeOf(err.Error()))
	fmt.Println(reflect.TypeOf(err))
}
