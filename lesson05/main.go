package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//生成模版,见hello.tmpl

	//解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("parse file failed,err:%v", err)
		return
	}

	//执行模版
	u1 := User{
		Name:   "阿斯顿",
		Age:    18,
		Gender: "man",
	}
	m1 := map[string]interface{}{
		"Name":   "asas",
		"Age":    18,
		"Gender": "as",
	}
	Hobby := []string{
		"lan",
		"zu",
		"shuang",
	}
	all1 := map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": Hobby,
	}
	err = t.Execute(w, all1)
	if err != nil {
		fmt.Println("execute file failed,err:%v", err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("start server failed,err:%v", err)
	}
	return

}
