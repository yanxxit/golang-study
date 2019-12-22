package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func post()  {
	url := `http://httpbin.org/post?key=123`

	// 填充表单，类似于net/url
	args := &fasthttp.Args{}
	args.Add("name", "test")
	args.Add("age", "18")

	status, resp, err := fasthttp.Post(nil, url, args)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println(string(resp))
}

func main() {
	post()
	url := `http://httpbin.org/get`

	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}
	fmt.Println(status)
	if status != fasthttp.StatusOK {
		fmt.Println("请求没有成功:", status)
		return
	}

	fmt.Println(string(resp))
}
