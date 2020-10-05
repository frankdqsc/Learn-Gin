package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//GET 请求 URL ? 后面的是 querystring 参数
//key= value格式，多个key-value用 &连接

func main(){
	r:= gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器发请求携带的 query string 参数
		//方式1：
		name := c.Query("query") //通过query 获取请求中携带 querystring 参数
		age := c.Query("age")
		//方式2：
		//name:= c.DefaultQuery("query","somebody")
		//方式3：
		//name,ok := c.GetQuery("query") //取不到第二个参数返回false
		//if !ok{
		//	//取不到
		//	name = "someday"
		//}
		c.JSON(http.StatusOK,gin.H {
			"name":name,
			"age": age,
		})
	})

	err := r.Run(":9090") //启动 server
	if err!= nil{
		fmt.Println("server failed")
	}
}