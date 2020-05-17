# email

发送邮件是一个很常见的需求：用户邮箱验证、邮箱召回等。Go 语言标准库自带 net/smtp 库，实现了 smtp 协议，用于发送邮件。然而这个库比较原始，使用不方便，而且官方声明不再增加新功能。于是乎出现了一些强大的第三方邮件库，今天推荐的这个库就是其中之一。

项目地址：https://github.com/jordan-wright/email，Star 数 1.3k+。

## 简介

email 包的设计易于使用，但又足够灵活以免受到限制。目的是为开发者提供友好的电子邮件接口。

该包当前支持以下功能：

- From, To, Bcc, 和 Cc；
- 邮件地址同时支持 "test@example.com" 和 "First Last test@example.com” 两种形式；
- 正文支持普通文本和 HTML；
- 附件支持；
- 已读回馈；
- 自定义协议头；
- 。。。

## 快速使用

先安装：

```
$ go get github.com/jordan-wright/email
```

后使用，使用 Gmail 发送邮件：

```
e := email.NewEmail()
e.From = "Jordan Wright <test@gmail.com>"
e.To = []string{"test@example.com"}
e.Bcc = []string{"test_bcc@example.com"}
e.Cc = []string{"test_cc@example.com"}
e.Subject = "Awesome Subject"
e.Text = []byte("Text Body is, of course, supported!")
e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
```

- 通过 NewEmail 获取一个 Email 的实例；
- From 设置发件人；
- To 设置发给谁，支持多人；
- Cc 设置抄送人，支持多人；
- Bcc 设置密抄人，支持多人；
- Subject 指定邮件标题；
- Text 指定普通文本邮件正文；
- HTML 指定 HTML 格式邮件正文；
- 调用 Email 实例的 Send 方法进行邮件发送：第一个参数是 smtp 地址，第二个参数是 smtp.Auth 实例；

可见该库最终还是使用标准库 net/smtp 进行实际的邮件发送。

### 注意事项

- From 中如果包含中文，比如：

```
e.From = "Go语言中文网<polaris@studygolang.com>”
```

收到的邮件不会显示发件人，有人提了一个 PR 支持 non-US-ASCII，但还未被 merge；根据该 PR 的提示，我们可以这样解决此问题：

```
e.From = mime.QEncoding.Encode("UTF-8", "Go语言中文网") + "<polaris@studygolang.com>"
```

- 如果同时指定 Text 和 HTML，则 Text 的内容会被忽略；
- 虽然默认情况下会尝试切换到 TLS，但如果想显示使用 TLS，则调用 Email#SendWithTLS 方法，具体代码如下：

```
tlsConfig := &tls.Config{
 InsecureSkipVerify: true,
  ServerName:         smtpHost,
}
e.SendWithTLS(addr, auth, tlsConfig)
```

- 因为 Email 结构体的字段都是导出的，因此可以通过 &email.Email{} 这种形式创建实例，并直接为各个字段赋值；

## 带附件的邮件

如果邮件中需要带上附件，这个库很方便：

```
e.AttachFile("test.txt")
```

也可以是通过 io.Reader 读取文件：

```
func (e *Email) Attach(r io.Reader, filename string, c string) (a *Attachment, err error)
```

其中 filename 是放入附件显示的文件名，c 是 Content-Type，如果为空，则使用 application/octet-stream。

## 重用连接

该库还支持连接池，例如：

```
p, err := email.NewPool(
  "smtp.qq.com:25",
  4,
  smtp.PlainAuth("", "274768166@qq.com", "password", "smtp.qq.com"))
```

- 第 2 个参数 4 表示最多开启 4 个连接

通过 p 的 Send 方法执行发送操作：

```
p.Send(e, 10e9)
```

完整例子如下：

```
package main
import (
 "log"
 "mime"
 "net/smtp"
 "net/textproto"
 "strconv"
 "sync"
 "github.com/jordan-wright/email"
)
func main() {
 var wg sync.WaitGroup
 ch := make(chan *email.Email, 4)
 err := sendEmailByPool(ch, &wg)
 if err != nil {
  panic(err)
 }
 for i := 0; i < 4; i++ {
  wg.Add(1)
  ch <- &email.Email{
   To:      []string{"polaris@studygolang.com"},
   From:    mime.QEncoding.Encode("UTF-8", "Go语言中文网") + "<274768166@qq.com>",
   Subject: "Pool" + strconv.Itoa(i),
   HTML:    []byte("<h1>这是 HTML 正文</h1>"),
   Headers: textproto.MIMEHeader{},
  }
 }
 wg.Wait()
 close(ch)
}
func sendEmailByPool(ch <-chan *email.Email, wg *sync.WaitGroup) error {
 p, err := email.NewPool(
  "smtp.qq.com:25",
  4,
  smtp.PlainAuth("", "274768166@qq.com", "password", "smtp.qq.com"))
 if err != nil {
  return err
 }
 for i := 0; i < 4; i++ {
  go func() {
   for e := range ch {
    err := p.Send(e, 10e9)
    if err != nil {
     log.Println("Send Email fail, err:", err)
    } else {
     log.Println("Send Email Successfully!")
    }
    wg.Done()
   }
  }()
 }
 return nil
}
```

## 总结

通过上面的介绍，你应该掌握了该库的使用，该库是不是满足了你对发邮件的需求？

最后附上一个完整的带附件的发邮件程序：

```
package main
import (
 "log"
 "mime"
 "net/smtp"
 "strings"
 "github.com/jordan-wright/email"
)
func main() {
 sendEmail("测试第三方 email 库", "xuxinhua@studygolang.com")
}
func sendEmail(subject string, tos ...string) error {
 e := email.NewEmail()
 smtpUsername := "274768166@qq.com"
 e.From = mime.QEncoding.Encode("UTF-8", "Go语言中文网") + "<274768166@qq.com>"
 e.To = tos
 e.Subject = subject
 e.HTML = []byte("<h1>HTML 正文</h1>")
 e.AttachFile("zap.log")
 auth := smtp.PlainAuth("", smtpUsername, "password", "smtp.qq.com")
 err := e.Send("smtp.qq.com:25", auth)
 if err != nil {
  log.Println("Send Mail to", strings.Join(tos, ","), "error:", err)
  return err
 }
 log.Println("Send Mail to", strings.Join(tos, ","), "Successfully")
 return nil
}
```

- 测试时注意将发件人、收件人和密码改为你自己的。