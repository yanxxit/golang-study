package main

import (
	"fmt"
	gomail "gopkg.in/gomail.v2"
)

var email_html = `Hello <a href = \"http://blog.csdn.net/liang19890820\">统计结果</a>`

func main() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "yanxx@infoepoch.com", "闫肖肖") // 发件人
	m.SetHeader("To",                                        // 收件人
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
	m.SetBody("text/html", email_html) // 正文
	// m.SetBody("我是正文") // 正文
	m.Attach("./test.txt") //添加附件

	d := gomail.NewPlainDialer("smtp.exmail.qq.com", 465, "yanxx@infoepoch.com", "Qinfen86") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	//文件压缩
	fmt.Println("发送成功！")
}
