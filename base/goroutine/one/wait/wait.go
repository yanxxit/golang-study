package wait

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.golang.org/",
	"http://www.google.com/",
	"http://www.qq.com/",
}

/**
aaaaa
 */
func GetMore() {
	for _, url := range urls {
		//每个url启动一个goruntine，同时给wg加1
		wg.Add(1)
		//Lannch a goruntine to fetch the url.
		go func(url string) {
			//当前goroutinoe 结束后给wg计数减1，wg.Done()等价于wg.Add(-1)
			//defer wg.Add(-1)
			defer wg.Done()
			//发送HTTP get请求并打印HTTP返回码
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)
			}
		}(url)
	}
	//等待所有请求结束
	wg.Wait()
}
