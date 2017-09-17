package crypto

import (
	"fmt"
	"testing"
	"encoding/base64"
)

func TestAesStringEncrypt(t *testing.T) {
	res, err := AesStringEncrypt("oKXUCj1MOddnp-sCpGi1J1dg3TyM", "wechat-mobile-17")
	fmt.Println("加密：", res, err) //ybXYkem1X9drSV9Bs1iN3YNeSjKY/SvnjOZYwHi3qdo=
}

func TestAesStringDecrypt(t *testing.T) {
	res, err := AesStringDecrypt("ybXYkem1X9drSV9Bs1iN3YNeSjKY/SvnjOZYwHi3qdo=", "wechat-mobile-17")

	fmt.Println("解密：", res, err) //ybXYkem1X9drSV9Bs1iN3YNeSjKY/SvnjOZYwHi3qdo=
}

func TestAes(t *testing.T) {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("wechat-mobile-17")
	result, err := AesEncrypt([]byte("oKXUCj1MOddnp-sCpGi1J1dg3TyM"), key)
	if err != nil {
		panic(err)
	}
	jm := base64.StdEncoding.EncodeToString(result);
	jiemi, _ := base64.StdEncoding.DecodeString(jm);
	origData, err := AesDecrypt(jiemi, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
