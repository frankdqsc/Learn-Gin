package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:= gin.Default()
	r.LoadHTMLFiles("./login.html","./index.html")
	r.GET("/login",func(c *gin.Context){
		c.HTML(http.StatusOK,"login.html",nil)
	})

	// login post
	r.POST("/login", func(c *gin.Context){
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":username,
			"Password": password,
		})
	})

	err := r.Run(":9090") //启动 server
	if err!= nil{
		fmt.Println("server failed")
	}
}