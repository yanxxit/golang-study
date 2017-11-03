package main

import (
	"os"
	"io"
	"bytes"
	"fmt"
	"bufio"
)

//Usage ./log_find aaaa.log keyword
func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("invalid args")
	}
	f, err := os.OpenFile(args[1], os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	in := bufio.NewReader(f)
	i := 0
	for true {
		line, isPrefix, err := in.ReadLine()
		if !isPrefix {
			line = append(line, byte('\n'))
		}
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println(err)
			return
		}
		if !bytes.Contains(line, []byte(args[2])) {
			continue
		}
		out, err := os.OpenFile(fmt.Sprint(i, ".log"), os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			fmt.Println(err)
			return
		}
		j := 0
		for j < 2<<20 {
			n, err := out.Write(line)
			if err != nil {
				fmt.Println(err)
				return
			}
			j += n
			line, isPrefix, err = in.ReadLine()
			if !isPrefix {
				line = append(line, byte('\n'))
			}
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Println(err)
				return
			}
		}
		i++
	}
}
