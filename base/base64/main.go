package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	input := []byte("hello golang base64 快乐编程http://www.01happy.com +~")

	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(string(decodeBytes))

	// 如果要用在url中，需要使用URLEncoding  针对数字 解析的结果 出现 + 就会解析失败。
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Println(uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(uDec))
}
