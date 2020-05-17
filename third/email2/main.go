package main

import (
	"fmt"
	"mime"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	fmt.Println("admin")
	e := email.NewEmail()
	from := mime.QEncoding.Encode("UTF-8", "果果")
	e.From = from + "bot@tingo66.com" // From 设置发件人；
	// fmt.Println(e.From)
	e.To = []string{"yanxxit@163.com"} // To 设置发给谁，支持多人；
	// e.Bcc = []string{"test_bcc@example.com"}                  //Bcc 设置密抄人，支持多人；
	// e.Cc = []string{"test_cc@example.com"}                    //Cc 设置抄送人，支持多人；
	e.Subject = "Awesome Subject" // Subject 指定邮件标题；
	// e.Text = []byte("Text Body is, of course, supported!") // Text 指定普通文本邮件正文；
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>") //HTML 指定 HTML 格式邮件正文；
	var err = e.Send("smtp.exmail.qq.com:465", smtp.PlainAuth("", "bot@tingo66.com", "A100s200", "smtp.exmail.qq.com"))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("end") // 尚未实现
}
