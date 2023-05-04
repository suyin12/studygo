package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func tempDemo(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	kua := func(name string) (string, error) {
		return name + "are a good man", nil
	}

	t := template.New("index.tmpl")
	t.Funcs(template.FuncMap{
		"kua": kua,
	})
	t, err := t.ParseFiles("./index.tmpl", "./ul.tmpl")
	//渲染模板
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	t.Execute(w, "小王子")
}

func main() {
	http.HandleFunc("/tempDemo", tempDemo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server failed,err:%v", err)
		return
	}
}
