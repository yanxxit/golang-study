package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var nickname = "闫肖肖" //加密
	nickname = base64.StdEncoding.EncodeToString([]byte(nickname))
	fmt.Println(nickname)
	//解密
	nick_name_baty, err := base64.StdEncoding.DecodeString(nickname)
	if err == nil {
		fmt.Println(string(nick_name_baty))
	}
	fmt.Println("还可以")
}
