package main

import (
	"encoding/json"
	"fmt"
	"mohoo-activity-bigwheel/models/struct/BindStruct"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func StrToInterface(){
	fmt.Println("interface 接口解析")
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	json.Unmarshal(b, &f)
	fmt.Println(f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	fmt.Println("----------------------------------------ok")
}

func createNewNulInterface()  {
	fmt.Println("闯将一个空的interface对象，进行解析JSON")
	fmt.Println("interface 赋值")
	var f interface{}
	fmt.Println(f)
	f = map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	fmt.Println(f)
	fmt.Println("-----------------------------------------------")
}

func init() {
	StrToInterface()
	createNewNulInterface();
	fmt.Println("打印数据：")
}

func main() {
	fmt.Println("通过结构体进行解析")
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Servers)
	fmt.Println(s.Servers[0].ServerName)

	binds();
}

func binds() {
	fmt.Println("解析绑定列表")
	var s BindStruct.BindData
	str := `{"flag":0,"result":{"openid":"oKXUCj1MOddnp-sCpGi1J1dg3TyM","nickname":"\"\\u95eb\\u6f47\"","headimg":"http://wx.qlogo.cn/mmopen/STan0NGQK3EWt4Ctfiab16jjRrg8V0O4LmC0Rcmk5XyG4SwI67UWY03wxnmn7iaibuEHjSbX1ibWUdibMICcDG3iaiaxxBGDS0Nx4Fia/0","accounts":[{"deviceno":"17721021494","devicename":"手机","bindtime":"2017-05-08 15:29:38","bindtdate":"2016-11-01 14:09:03","bpn":"97594334564","sort":"100"},{"deviceno":"18918901568","devicename":"手机","bindtime":"2017-04-24 08:19:41","bindtdate":"2017-04-24 08:19:41","bpn":"66703457004","sort":"1"},{"deviceno":"18916777768","devicename":"手机","bindtime":"2017-03-09 09:15:26","bindtdate":"2016-09-19 19:56:35","bpn":"42058671585","sort":"1"},{"deviceno":"18916491263","devicename":"手机","bindtime":"2017-03-09 09:16:51","bindtdate":"2016-09-19 19:56:41","bpn":"46930468171","sort":"1"},{"deviceno":"KD1108346442","devicename":"宽带","bindtime":"2017-05-09 11:13:40","bindtdate":"2017-05-09 11:13:40","bpn":null,"sort":"1"},{"deviceno":"18930840402","devicename":"手机","bindtime":"2017-04-22 09:33:39","bindtdate":"2017-04-22 09:33:39","bpn":"92586218453","sort":"1"},{"deviceno":"18016038006","devicename":"手机","bindtime":"2017-03-08 21:29:07","bindtdate":"2016-09-26 22:44:31","bpn":"84936470609","sort":"1"},{"deviceno":"ZH10039609","devicename":"宽带","bindtime":"2017-04-21 18:29:25","bindtdate":"2017-04-21 18:29:25","bpn":"96846807184","sort":"1"},{"deviceno":"KD1005497439","devicename":"宽带","bindtime":"2017-05-02 13:24:25","bindtdate":"2016-10-18 15:23:53","bpn":"84936470609","sort":"1"},{"deviceno":"17301770123","devicename":"手机","bindtime":"2017-04-30 16:13:00","bindtdate":"2017-04-30 16:13:00","bpn":"98976080301","sort":"1"},{"deviceno":"KD1040808907","devicename":"宽带","bindtime":"2017-05-02 13:35:33","bindtdate":"2017-05-02 13:35:33","bpn":"93597755509","sort":"1"},{"deviceno":"KD1108349849","devicename":"宽带","bindtime":"2017-05-09 11:05:56","bindtdate":"2017-05-09 11:05:56","bpn":null,"sort":"1"},{"deviceno":"KD1107030463","devicename":"宽带","bindtime":"2017-04-07 09:24:03","bindtdate":"2017-04-07 09:24:03","bpn":null,"sort":"1"},{"deviceno":"AD0004131592","devicename":"宽带","bindtime":"2017-03-24 15:22:59","bindtdate":"2017-03-24 15:22:59","bpn":"26140018902","sort":"1"},{"deviceno":"15317328043","devicename":"手机","bindtime":"2017-05-04 09:41:11","bindtdate":"2017-05-04 09:41:11","bpn":"74260610932","sort":"1"},{"deviceno":"18918803993","devicename":"手机","bindtime":"2017-04-19 15:04:53","bindtdate":"2017-04-19 15:04:53","bpn":"46930468171","sort":"1"},{"deviceno":"KD1108348623","devicename":"宽带","bindtime":"2017-05-09 11:12:16","bindtdate":"2017-05-09 11:12:16","bpn":null,"sort":"1"}]}}`

	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
	fmt.Println(s.Result.Openid)
	fmt.Println(s.Result.Headimg)
	fmt.Println(s.Result.Accounts)
	fmt.Println("使用解析表")
	fmt.Println(s.Result.Accounts[0])
}
