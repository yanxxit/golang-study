package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5Encode(value string) string {
	data := []byte(value)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str1
}

func MD5Encode2(value string) string {
	w := md5.New()
	io.WriteString(w, value)                 //将str写入到w中
	md5str1 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5str1
}

func main() {
	str := "abc123"

	fmt.Println(MD5Encode(str))

	//方法二
	fmt.Println(MD5Encode2(str))
}
