package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//生成模版f.tmpl

	//解析模版
	t := template.New("f.tmpl") //创建模版名字f，名字一定要与模版的名字对应上
	//告诉模版引擎，多个一个自定义模版函数kua
	kua := func(name string) (string, error) {
		return name + "你真帅", nil
	}
	t.Funcs(template.FuncMap{
		"kua99": kua,
	})
	_, err := t.ParseFiles("f.tmpl")
	if err != nil {
		fmt.Printf("parse file failed,err:%v\n", err)
		return
	}
	//渲染模版
	name := "小王子"
	t.Execute(w, name)

}

func f2(w http.ResponseWriter, r *http.Request) {
	//生成模版
	//解析模版
	t, err := template.ParseFiles("t.tmpl", "u1.tmpl")
	if err != nil {
		return
	}
	name := "s"
	t.Execute(w, name)
	//执行模版
}
func main() {
	
	http.HandleFunc("/", f1)
	http.HandleFunc("/demo", f2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("listen server failed,err:%v\n", err)
	}

}
