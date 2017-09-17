package httputil

import (
	"fmt"
	"testing"
	//"encoding/json"
	//"github.com/astaxie/beego/httplib"
	//"time"
	//"github.com/astaxie/beego/logs"
)

func TestGet(t *testing.T) {
	body, terr := Get("http://m.sh.189.cn/service/node/crypto?data=oKXUCj1MOddnp-sCpGi1J1dg3TyM,abc&key=wechat-mobile-201604&type=0");
	fmt.Println(terr)
	fmt.Println("get 请求结果" + string(body))
}

//
//func TestPost(t *testing.T) {
//	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
//	var f interface{}
//	err := json.Unmarshal(b, &f)
//	if err == nil {
//		fmt.Println(f)
//	}
//
//	req, err := httplib.Post("http://127.0.0.1:6200/v1/proxy/").SetTimeout(5*time.Second, 5*time.Second).JSONBody(f)
//	body, _ := req.Bytes()
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Post 请求结果" + string(body))
//}

//
//func TestPostForm(t *testing.T) {
//	var PayParams map[string]interface{}
//	PayParams = make(map[string]interface{});
//	PayParams["mobile"] = "17721021494"
//	PayParams["channel"] = "2"
//	PayParams["desKey"] = "dzqd-wt-flow"
//
//	body, terr := PostForm("http://172.16.50.138:8091/csb/1.0/Encrypt", PayParams);
//	if terr != nil {
//
//	}
//	fmt.Println("PostForm 请求结果" + string(body))
//
//	//str := `{"mobile":"17721021494","channel":"2","desKey":"dzqd-wt-flow"}`
//	//PayParams := make(map[string]interface{});
//	//json.Unmarshal([]byte(str), &PayParams)
//	//
//	//req := httplib.Post("http://172.16.50.138:8091/csb/1.0/Encrypt").SetTimeout(5*time.Second, 5*time.Second)
//	//for k, v := range PayParams {
//	//	req.Param(k, ComUtil.TransInterfaceToString(v))
//	//}
//	//data, err := req.Bytes()
//	//logs.Info("request-postForm-data: "+string(data), err)
//	//fmt.Println(PayParams)
//
//}
//
//func TestPut(t *testing.T) {
//	var PayParams map[string]interface{}
//	PayParams = make(map[string]interface{});
//	PayParams["contact_id_card"] = "310115198907222517"
//	PayParams["contact_name"] = "叶涛"
//	PayParams["contact_phone"] = "18117562317"
//	PayParams["into_park_time"] = "2017-07 -14"
//
//	//body, terr := Put("http://httpbin.org/put", PayParams,"");
//	body, terr := Put("http://172.16.50.141:8080/v1/disney_pay/draw_prize/57357/oKXUCjzkP8KjkFdobdFIuhxK5dGw", PayParams);
//	if terr != nil {
//
//	}
//	fmt.Println("put 请求结果" + string(body))
//}
