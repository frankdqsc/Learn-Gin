package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)
//遇事不决 先写注释

//静态文件 html 页面上用到的样式文件 .css js文件 图片

func main(){
	//模板定义
	r := gin.Default()
	//加载静态文件
	r.Static("/xxx","./statics")
	//gin 框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe":func(str string)template.HTML{
			return template.HTML(str)
		},
	})
	//r.LoadHTMLFiles("templates/index.tmpl","templates/users/index.tmpl")  //模板解析  多个文件时太麻烦
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context){
		//HTTP请求
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{ //模板渲染
			"title":"<a href = 'www.google.com'>google.com</a>",
		})
	})

	r.GET("/users/index",func(c *gin.Context){
		//HTTP请求
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{ //模板渲染
			"title":"users",
		})
	})

	//返回从网上下载得模板
	r.GET("/home", func(c *gin.Context){
		//没有给模板起名字 则使用默认的文件名 home.html
		c.HTML(http.StatusOK,"home.html",nil)  //没有返回渲染得数据返回  nil
	})

	err := r.Run(":9090") //启动 server
	if err!= nil{
		fmt.Println("server failed")
	}
}