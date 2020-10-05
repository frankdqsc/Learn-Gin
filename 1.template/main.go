package main

import(
	"fmt"
	"net/http"
	"html/template"
)
//遇事不决 先写注释
type User struct{
	Name string
	Gender string
	Age int
}
func sayHello(w http.ResponseWriter, r *http.Request){

	//2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")  //注意 go build位置
	if err!= nil{
		fmt.Println("Parse template failed ,err:%v", err)
		return
	}
	//3.渲染模板

	u1 := User{  //结构体首字母大写
		Name: "小王子111",
		Gender:"男",
		Age: 10,
	}
	m1:= map[string]interface{}{  //map 首字母没必要大写
		"name":"小王子",
		"gender":"男",
		"age": 18,
	}

	hobbyList := []string{
		"basketball",
		"soccer",
		"twocolorball",
	}
	t.Execute(w,map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby":hobbyList,
	})
}
func main(){
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":9000",nil)
	if err!= nil{
		fmt.Println("Http server start failed, err:%v",err)
		return
	}
}