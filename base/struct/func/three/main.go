package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

func (c *Car) Set(name string, age int) {
	c.Name = name
	c.Age = age
}

type Car2 struct {
	Name string
}

//Go有匿名字段特性
type Train struct {
	Car
	Car2
	createTime time.Time
	//count int   正常写法，Go的特性可以写成
	int
}

//给Train加方法，t指定接受变量的名字，变量可以叫this，t，p
func (t *Train) Set(age int) {
	t.int = age
}

func main() {
	var train Train
	train.int = 300 //这里用的匿名字段写法，给Age赋值
	//(&train).Set(1000)
	train.Car.Set("huas", 100)
	train.Car.Name = "test" //这里Name必须得指定结构体
	fmt.Println(train)
}
