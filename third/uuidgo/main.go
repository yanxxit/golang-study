package main

import uuid "github.com/satori/go.uuid"

func main() {
	// 创建 UUID v4
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	println(`生成的UUID v4：`)
	println(u1.String())

	//uuid.new

	// 创建可以进行错误处理的 UUID v4
	u2 := uuid.NewV4()

	println(`生成的UUID v4：`)
	println(u2.String())

	// 解析 字符串 到 UUID
	u2, err2 := uuid.FromString(`6ba7b810-9dad-11d1-80b4-00c04fd430c8`)
	if err2 != nil {
		println(`解析 字符串 到 UUID 时出错`)
		panic(err2)
	}
	println(`解析 字符串 到 UUID 成功！解析到的 UUID 如下：`)
	println(u2.String())
}
