package main

import "net/url"

/**

var options = {
       url: url,
        method: 'POST',
        json: true,
        body: params,
        headers: headers,
        timeout: apis.timeout
    };

*/

type Request struct {
	Url     string
	Method  string
	Json    bool
	Body    map[string]interface{}
	Headers map[string]interface{}
	TimeOut int
}

func (h *Request) Get(url string) (body interface{}, err error) {
	return body, err
}

func main() {
	h := &Request{Method: "Get"}
	h.Get("")
}
