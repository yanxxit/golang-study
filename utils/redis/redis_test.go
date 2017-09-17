package redis

import (
	"testing"
	//"time"
	"fmt"
)

func TestGetCache(t *testing.T) {
	data := GetCache("177210214942")
	if data == "" {
		fmt.Println("未获取数据")
	} else {
		fmt.Println("以获取数据")
	}
	fmt.Println("--", data)
}

//func TestAddCache(t *testing.T) {
//	NameSpace = "shdx:"
//	expires := 10000 * time.Second
//	AddCatch("18016038006", `{"data":"123"}`, expires)
//
//	NameSpace = "shdx:name:"
//	AddCatch("18016038006", `{"data":"123"}`, expires)
//}
