package restful

import (
	"testing"
	"fmt"
)

//func TestParams(t *testing.T) {
//	var data map[string]interface{}
//	fmt.Println(Params(1, data, ""))
//}

func TestFailCode(t *testing.T) {
	fmt.Println(FailCode("1", "错误结果"))
	//fmt.Println(Fail( "错误结果"))
}
