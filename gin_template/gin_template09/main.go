package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t := template.New("index.tmpl").Delims("{[", "]}")
	t, err := t.ParseFiles("./index.tmpl")
	//渲染模板
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	t.Execute(w, "小王子")
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server failed,err:%v", err)
		return
	}
}
