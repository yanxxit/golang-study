package restful

type ErrorCode struct {
	Status int `db:"status" json:"status"`
	Data   string `db:"data" json:"data"`
}

type PhpErrorStatusCode struct {
	Status int `db:"status" json:"status"`
}

type GoErrorCode struct {
	ErrorCode int `db:"error_code" json:"error_code"`
	ErrorMsg  string `db:"error_msg" json:"error_msg"`
}

//代理请求通用入参
type ProxyPrams struct {
	URL    string `db:"url" json:"url"`            //请求地址
	Params interface{} `db:"params" json:"params"` //请求入参
}

type CsbResCode struct {
	ResData interface{} `db:"res_data" json:"resData"`
	ResMsg  string `db:"res_msg" json:"resMsg"`
	ResNum  int `db:"res_num" json:"resNum"`
}
