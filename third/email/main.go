package main

import (
	"fmt"
	"io/ioutil"

	gomail "gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
)

type semail struct {
	From     string `yaml:"from"`
	Pwd      string `yaml:"pwd"`
	FromName string `yaml:"from_name"`
}

type conf struct {
	Host  string `yaml:"host"`
	User  string `yaml:"user"`
	Pwd   string `yaml:"pwd"`
	Email semail
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("../../dev.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

var emailHTML = `Hello <a href = \"http://blog.csdn.net/liang19890820\">统计结果</a>`

func main() {
	var c conf
	conf := c.getConf()
	fmt.Println(conf.Email.FromName)
	fmt.Println(conf.Email.Pwd)
	m := gomail.NewMessage()
	m.SetAddressHeader("From", conf.Email.From, conf.Email.FromName) // 发件人
	m.SetHeader("To",                                                // 收件人
		m.FormatAddress("yanxxit@163.com", "乔峰"),
		// m.FormatAddress("********@qq.com", "郭靖"),
	)

	//抄送
	// m.SetHeader("Cc",
	// 	m.FormatAddress("xxxx@foxmail.com", "收件人"),
	// )
	// m.SetHeader("Bcc",
	// 	m.FormatAddress("xxxx@gmail.com", "收件人")) //暗送
	m.SetHeader("Subject", "活动日志数据统计") //设置主题
	m.SetBody("text/html", emailHTML)  // 正文
	// m.SetBody("我是正文") // 正文
	m.Attach("./Gomail.md") //添加附件

	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, conf.Email.From, conf.Email.Pwd) // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	//文件压缩
	fmt.Println("发送成功！")
}
