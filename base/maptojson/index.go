package main

import (
	"fmt"
	"encoding/json"
)

type MyStruct struct {
	Admin string `json "admin"`
	Pwd   string `json "pwd"`
}

func main() {
	var fail map[string]interface{} = make(map[string]interface{})

	//map
	var t map[string]interface{} = make(map[string]interface{})
	t["age"] = 12
	t["name"] = "inter"

	//struct
	m := new(MyStruct)
	m.Admin = "admin"
	m.Pwd = "12312312"

	n := new(MyStruct)
	n.Admin = "admin111"
	n.Pwd = "111111111"

	//列表
	//source := []string{"apple", "orange", "plum", "banana", "grape"}
	//list := []interface{}{m, n}

	fail["status"] = 0
	fail["info"] = ""
	fail["data"] = t
	fail["detail"] = m
	//fail["list"] = source
	//fail["list2"] = list
	//fmt.Println("map格式：", string(fail))
	fmt.Println("map格式：", fail)

	content, _ := json.MarshalIndent(fail, "", "  ")
	fmt.Println("JSON格式数据：" + string(content))

	b, _ := json.Marshal(fail)
	fmt.Println("JSON.stringify:" + string(b))
}
