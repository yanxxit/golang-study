package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ain() {

}

//转换的思路 入参s是interface 主要针对string interface map 等 然后结构话map[string]interface{}

func TransInterfaceToMapList(data interface{}) (resData map[string]interface{}, err error) {
	mybyte, err := TransInterfaceToByte(data)
	resData = make(map[string]interface{})
	var body map[string]interface{}
	err = json.Unmarshal(mybyte, &body)
	for k, v := range body {
		fmt.Println("============= 遍历map", k, v, reflect.TypeOf(v)) //未获取到值，name ==nil
		resData[k] = v
	}
	return resData, err
}

func TransInterfaceToMapOne(data interface{}) (body map[string]interface{}, err error) {
	mybyte, err := TransInterfaceToByte(data)
	err = json.Unmarshal(mybyte, &body)
	return body, err
}

func TransInterfaceToByte(data interface{}) (mybyte []byte, err error) {
	switch v := data.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	default:
		return nil, errors.New("非string或者[]byte")
	}
}
