package wait

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.taobao.com/",
	"http://www.baidu.org/",
	"http://www.qq.com/",
	"https://cnodejs.org/api/v1/topic_collect/yanxxit",
}

type CnodeStruct struct {
	Success   string `json:"success"`
	Error_msg string `json:"error_msg"`
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
				fmt.Println("--->>>>", url, resp.Status)
				//程序在使用完回复后必须关闭回复的主题
				defer resp.Body.Close()
				if strings.Contains(url, "cnode") {
					var cnode CnodeStruct
					body, _ := ioutil.ReadAll(resp.Body)
					json.Unmarshal(body, cnode)
					fmt.Println("校验：", url, string(body))
					fmt.Println("----L>>", cnode.Error_msg)
					fmt.Println("----L>>", cnode.Success)
					fmt.Println("----L>>", cnode)
				}
			}
		}(url)
	}
	//等待所有请求结束
	wg.Wait()
}
