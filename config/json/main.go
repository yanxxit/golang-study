package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//目标
//读取配置文件，可以识别到开发，测试，生产环境，可以识别到通用的配置
//通用配置
type CommonConfig struct {
	Port int32
	Host string
}

//特殊配置
type configuration struct {
	CommonConfig
	Enabled bool
	Path    string
	Name    string
	Title   string
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var Config configuration
var common CommonConfig

func init() {
	var confFile string
	currentEnv := "dev"
	fmt.Printf(`环境：%s读取json中的配置`, currentEnv)
	if currentEnv == "prod" {
		confFile = "prod.conf.json"
	} else if currentEnv == "test" {
		confFile = "test.conf.json"
	} else {
		confFile = "dev.conf.json"
	}
	file, _ := os.Open(confFile)

	defer file.Close()

	isExit, err := PathExists("common.conf.json")
	if err == nil && isExit {
		//通用
		commonFile, _ := os.Open("common.conf.json")
		defer commonFile.Close()
		common_decoder := json.NewDecoder(commonFile)
		common_decoder.Decode(&Config)
	} else {
		fmt.Println("通用配置文件未读取")
	}
	//环境
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)

	if err != nil {
		fmt.Println("文件解析失败:", err)
	}

}

func main() {
	fmt.Println(Config.Path)
	fmt.Printf(`服务名称(%s): %t`, Config.Name, Config.Enabled)
	fmt.Println()
	fmt.Printf("服务地址：%s:%d", Config.Host, Config.Port)
}
