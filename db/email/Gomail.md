简述

Go 提供了一个 smtp（Simple Mail Transfer Protocol - 简单邮件传输协议）库作为其网络包的一部分，“net/smtp”公开了一些可立即使用的有用功能。

Gomail - 一个比较成熟的第三方包，提供了一个快速、简单的解决方案，可以轻松地发送电子邮件。

简述
Gomail
特性
使用
常用邮箱
开启服务
Gomail

Gomail 是一个简单、高效的发送电子邮件包，它经过良好的测试和记录。

Gomail 只能使用 SMTP 服务器发送电子邮件，但是 API 比较灵活的，很容易实现其他方法使用本地 Postfix、API 等发送电子邮件。

项目地址：https://github.com/go-gomail/gomail
文档：https://godoc.org/gopkg.in/gomail.v2
示例：https://godoc.org/gopkg.in/gomail.v2#example-package
特性

Gomail 支持：

附件
嵌入图像
HTML 和文本模板
特殊字符的自动编码
SSL 和 TLS
使用相同的 SMTP 连接发送多封电子邮件
使用

下载 Gomail，解压缩至 $GOPATH\src（例如：E:\Works\Golang\src），新建并编写 sendMail.go：

```go
package main

import (
    "gomail"
)

func main() {
    m := gomail.NewMessage()
    m.SetAddressHeader("From", "550755606@qq.com", "一去、二三里")  // 发件人
    m.SetHeader("To",  // 收件人
        m.FormatAddress("********@163.com", "乔峰"),
        m.FormatAddress("********@qq.com", "郭靖"),
    )
    m.SetHeader("Subject", "Gomail")  // 主题
    m.SetBody("text/html", "Hello <a href = \"http://blog.csdn.net/liang19890820\">一去丶二三里</a>")  // 正文

    d := gomail.NewPlainDialer("smtp.qq.com", 465, "550755606@qq.com", "*********")  // 发送邮件服务器、端口、发件人账号、发件人密码
    if err := d.DialAndSend(m); err != nil {
        panic(err)
    }
}

```
运行：go run sendMail.go，发送成功之后，就能收到邮件了。

这里写图片描述

常用邮箱

列举一些常用的邮箱，可以用来测试：

QQ 邮箱 
POP3 服务器地址：qq.com（端口：995） 
SMTP 服务器地址：smtp.qq.com（端口：465/587）

163 邮箱： 
POP3 服务器地址：pop.163.com（端口：110） 
SMTP 服务器地址：smtp.163.com（端口：25）

126 邮箱： 
POP3 服务器地址：pop.126.com（端口：110） 
SMTP 服务器地址：smtp.126.com（端口：25）

开启服务

出于邮箱安全的考虑，很多邮箱缺省是关闭 POP3/SMTP 服务的，需要登录邮箱设置后开启。

以 QQ 邮箱为例，进入邮箱“设置”，在“帐户”项里就可找到“POP3/SMTP服务”的设置项，进行开启。

这里写图片描述

温馨提示：登录第三方客户端时，密码框请输入“授权码”进行验证。

也就是说，在使用 QQ 邮箱发送邮件的时候，需要使用授权码，而不是 QQ 密码！

## 参考
- [gomail-github](https://github.com/go-gomail/gomail)
- [Golang 使用gomail包发送邮件](http://blog.csdn.net/wj199395/article/details/75206501?utm_source=debugrun&utm_medium=referral)

- [golang发送邮件（抄送，暗送，附件）](http://blog.csdn.net/qq_30949367/article/details/71076193)