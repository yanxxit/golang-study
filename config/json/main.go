package main

import (
	"encoding/json"
	"fmt"
	"os"
)


//目标
//读取配置文件，可以识别到开发，测试，生产环境，可以识别到通用的配置

type configuration struct {
	Enabled bool
	Path    string
}
var conf  configuration
func init()  {
	currentEnv :="dev"
	fmt.Println("读取json中的配置")
	confFile := currentEnv+".conf.json"
	file, _ := os.Open(confFile)
	defer file.Close()
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

func main() {
	fmt.Println(conf.Path)
	fmt.Println(conf.Enabled)
}
