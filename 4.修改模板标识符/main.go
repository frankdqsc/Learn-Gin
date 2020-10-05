package main

import (
	"fmt"
	"html/template"
	"net/http"
)
//遇事不决 先写注释
type User struct{
	Name string
	Gender string
	Age int
}

func index(w http.ResponseWriter, r*http.Request){
	//定义模板
	//解析模板
	t,err := template.New("index.tmpl").
		Delims("{[","]}").
		ParseFiles("./index.tmpl")
	if err != nil{
		fmt.Printf("parse template failed ,err:%v\n",err)
		return
	}
	name:= "小王子2222"
	//渲染模板
	t.Execute(w,name)

}
func main(){
	http.HandleFunc("/index",index)
	err := http.ListenAndServe(":9000",nil)
	if err!= nil{
		fmt.Println("Http server start failed, err:%v",err)
		return
	}
}