package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
)

func httpPost() {
	resp, err := http.Post("http://172.16.50.141/mohoo-telecom-sms/Rest/ValidateCode/sendVerify",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
//func main()  {
//	httpPost();
//}
