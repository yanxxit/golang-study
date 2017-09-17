package regexp

import (
	"regexp"
	"fmt"
)

func IsMobile(mobile string) bool {
	var validID = regexp.MustCompile(`^[1][123456789]\d{9}$`)
	return validID.MatchString(mobile)
}

func IsOpenid(openid string) bool {
	var validID = regexp.MustCompile(`^[\x00-\x7F]{28}$`)
	return validID.MatchString(openid)
}

func IsEmail(str string) bool {
	var validID = regexp.MustCompile(`^[-_A-Za-z0-9]+@([_A-Za-z0-9]+\.)+[A-Za-z0-9]{2,3}$`)
	return validID.MatchString(str)
}

func FormatNumber(str string) string {
	return str[0:3] + "****" + str[7:len(str)]
}

//还有些问题
func FormatName(str string) string {
	a := str[0:1]
	b := len(str)
	c := ""
	for i := 1; i < b; i ++ {
		fmt.Println("11")
		c += "*"
	}
	return a + c
}
