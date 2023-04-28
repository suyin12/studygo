package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 一个函数包含两个参数响应，请求
func sayHello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")
	_, _ = fmt.Fprintln(w, string(b))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http serve failed err:%v\n", err)
		return
	}
}
