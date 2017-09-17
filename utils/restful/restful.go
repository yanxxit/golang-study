package restful

import (
	"encoding/json"
	"fmt"
)

func Params(status string, data map[string]interface{}, info string) (content interface{}, errs error) {
	str := ""
	var params interface{}
	err := json.Unmarshal([]byte(str), &params)
	return params, err
}

func Fail(desc string) (content interface{}) {
	str := `{
	  "data": "` + desc + `",
	  "info": 0,
	  "status": 0
	}`
	var params interface{}
	json.Unmarshal([]byte(str), &params)
	return params
}

func FailCode(code string, msg string) (content interface{}) {
	str := `{
	  "error_code": ` + code + `,
	  "error_msg": "` + msg + `"
	}`

	fmt.Println(str)
	var params interface{}
	json.Unmarshal([]byte(str), &params)
	return params
}

func ErrorDealData(data []byte, info string) interface{} {
	var resData interface{}
	var errcode ErrorCode
	if err := json.Unmarshal(data, &errcode); err == nil && errcode.Status == 0 {
		if errcode.Data == "" {
			resData = Fail(info)
		} else {
			resData = Fail(errcode.Data)
		}
	} else {
		fmt.Println(err)
		resData = Fail(info)
	}
	return resData
}
