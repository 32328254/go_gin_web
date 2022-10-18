package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f2(w http.ResponseWriter, r *http.Request) {
	//生成模版
	//解析模版
	//自定义模版解析引擎Delims("{[","]}")
	t,err := template.New("index.tmpl").Delims("{[","]}").ParseFiles("index.tmpl")
	if err != nil {
		return
	}
	name := "s"
	t.Execute(w, name)
	//执行模版
}


func main() {
	http.HandleFunc("/index", f2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("listen server failed,err:%v\n", err)
	}
}
