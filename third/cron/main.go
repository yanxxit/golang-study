package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
)

func main() {
	i := 0
	c := cron.New()
	fmt.Println("cron running:start")
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
		fmt.Println("cron running:", i)
	})
	c.Start()
}
