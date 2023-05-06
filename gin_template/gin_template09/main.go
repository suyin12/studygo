package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	//解析模板前自定义函数
	t := template.New("index.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).Delims("{[", "]}")
	t, err := t.ParseFiles("./index.tmpl")
	//渲染模板
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	str1 := "<script>alter(123)</script>"
	str2 := "<a href='https://www.google.com'>谷歌</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("HTTP server failed,err:%v", err)
		return
	}
}
