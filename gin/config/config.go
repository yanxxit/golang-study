package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//profile variables
type Yaml struct {
	Env     string `yaml:"env"`
	Port    int  `yaml:"port"`
	Mongodb string `yaml:"mongodb"`
}

var Conf Yaml

func LoadConfig(dir string, env string) {

	fmt.Println("config init Loading:" + dir + "/" + env + ".yaml")
	yamlFile, err := ioutil.ReadFile(dir + "/" + env + ".yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		fmt.Println(err.Error())
	}
}

var dir = flag.String("config", "./", "input your config")
var env = flag.String("env", "dev", "input your env")

func init() {
	fmt.Println("配置初始化：config")
	//LoadConfig("./config", "dev")
	flag.Parse()
	LoadConfig(*dir, *env)
	fmt.Println("env:", Conf.Env)
}
