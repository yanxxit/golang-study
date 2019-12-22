package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"log"
)

type User struct {
	Name string
	Age  int8
}
// 据说 json-iteator 是目前golang中对json格式数据处理最快的包(比官方json包快6倍)，好像是滴滴团队开源的，使用起来也非常方便，
// http://jsoniter.com/migrate-from-go-std.html
func main() {
	user := User{
		Name: "tanggu",
		Age:  18,
	}
	var jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary
	// 序列化
	data, err := jsoniter.Marshal(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// 反序列化
	var people User
	err = jsoniter.Unmarshal(data, &people)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
}
