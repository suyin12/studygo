package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse failed,err:%v", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	user := UserInfo{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "小王子",
		"age":    18,
		"gender": "男",
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	tmpl.Execute(w, map[string]interface{}{
		"user":  user,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	// err := http.ListenAndServe(":9090", nil)()
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server failed,err:%v", err)
		return
	}
}
