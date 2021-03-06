package main

import (
	"fmt"
	"html/template"
	"net/http"
)
func index(w http.ResponseWriter, r*http.Request){
	//定义模板
	//解析模板
	t,err := template.ParseFiles("./index.tmpl")
	if err != nil{
		fmt.Printf("parse template failed ,err:%v\n",err)
		return
	}
	msg := "xiaowangzi"
	//渲染模板
	t.Execute(w,msg)
}

func home(w http.ResponseWriter, r*http.Request){
	//定义模板
	//解析模板
	t,err := template.ParseFiles("./home.tmpl")
	if err != nil{
		fmt.Printf("parse template failed ,err:%v\n",err)
		return
	}
	msg := "xiaowangzi"
	//渲染模板
	t.Execute(w,msg)
}
func index2(w http.ResponseWriter, r*http.Request){
	//定义模板 （模板继承的方式）
	//解析模板
	t,err:=template.ParseFiles("./templates/base.tmpl","./templates/index2.tmpl")
	if err!= nil{
		fmt.Println("template parse failed, err:%v",err)
		return
	}
	//渲染模板
	name := "xiaowangzi"
	t.ExecuteTemplate(w,"index2.tmpl",name)
}
func home2(w http.ResponseWriter, r*http.Request){
	//定义模板 （模板继承的方式）
	//解析模板
	t,err:=template.ParseFiles("./templates/base.tmpl","./templates/home2.tmpl")
	if err!= nil{
		fmt.Println("template parse failed, err:%v",err)
		return
	}
	//渲染模板
	name := "laowang"
	t.ExecuteTemplate(w,"home2.tmpl",name) //ExecuteTemplate
}
func main(){
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)

	http.HandleFunc("/index2",index2)
	http.HandleFunc("/home2",home2)

	err := http.ListenAndServe(":9000",nil)
	if err!= nil{
		fmt.Println("Http server start failed, err:%v",err)
		return
	}
}