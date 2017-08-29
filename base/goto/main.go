package main

import (
	"fmt"
	"time"
	"encoding/json"
)

const result = `{"name":"zhangsan","data":{"result0":100,"result1":200,"result2":300,"result3":400,"result5":100}}`

type Result struct {
	Data map[string]int32 `json:"data"`
}

func init() {

}

func main() {
getNewData:
	for i := 0; i < 10; i++ {
		for {
			fmt.Println(i)
			continue getNewData
		}
	}

	StartCac();
	fmt.Println("程序启动")
}

func StartCac() {
	t1 := time.Now() // get current time

	fmt.Println(t1)
	//logic handlers
	for i := 0; i < 1000; i++ {
		fmt.Print("*")
	}
	elapsed := time.Since(t1) //Go计算运行的时间
	fmt.Println("\nApp elapsed: ", elapsed)

	var r Result
	if err := json.Unmarshal([]byte(result), &r); err != nil {
		panic(err)
	}

	statics := make(map[int32]int32)
	for _, value := range r.Data {
		statics[value] = statics[value] + 1
	}

	fmt.Println(statics)
}
