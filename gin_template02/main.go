package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("Parse failed,err:%v", err)
		return
	}
	name := "小王子"
	tmpl.Execute(w, name)
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
