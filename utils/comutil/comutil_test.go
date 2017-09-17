package comutil

import (
	"testing"
	"fmt"
)

//func TestGetPayid(t *testing.T) {
//	for i := 0; i < 10; i++ {
//		fmt.Println(GetPayid())//长度无法确认
//	}
//}
var myJSON = `{
    "data": {
        "attachamount": "",
        "curtype": "RMB",
        "encodetype": "1",
        "orderamount": "10000",
    },
    "status": 200
}`

type ErrorCode struct {
	Status int `db:"status" json:"status"`
	Data   interface{} `db:"data" json:"data"`
}

func TestTransInterfaceToString(t *testing.T) {
	fmt.Println("int:" + TransInterfaceToString(1))
	fmt.Println("string" + TransInterfaceToString(`{"root":"mohoo"}`))
	mybyte := []byte(myJSON)
	fmt.Println("byte:", mybyte)
	fmt.Println("string" + TransInterfaceToString(mybyte))

	//map
	params := make(map[string]interface{})
	params["mobile"] = "15806111230"
	params["tkey"] = ""
	params["from"] = "iphone8"
	fmt.Println("map:" + TransInterfaceToString(params))

	//struct 结构体
	myStruct := new(ErrorCode)
	myStruct.Status = 100
	myStruct.Data = "成功"
	fmt.Println("struct:" + TransInterfaceToString(myStruct))

	//struct +  map 或者interface
	myStruct2 := new(ErrorCode)
	myStruct2.Status = 200
	myStruct2.Data = params
	fmt.Println("struct:" + TransInterfaceToString(myStruct2))
}
