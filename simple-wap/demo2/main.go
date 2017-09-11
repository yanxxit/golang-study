package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.ListenAndServe(":8083", nil)
}

//访问方式：http://127.0.0.1:8083/
