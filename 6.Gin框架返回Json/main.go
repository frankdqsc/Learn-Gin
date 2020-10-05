package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	r:= gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1：使用map
		//data := map[string]interface{}{
		//	"name":"小王子",
		//	"message":"hello world",
		//	"age": 18,
		//}
		//方法2：
		data:= gin.H{"name":"小王","message":"hello world","age": 18}

		c.JSON(http.StatusOK,data)
	})

	//方法3：结构体 灵活使用 tag 来对结构体字段做定制化操作
	type msg struct{
		Name string  `json:"name"`//结构体字段不可 首字母小写，小写可通过 tag
		Message string
		Age int
	}
	r.GET("/another_json", func(c *gin.Context) {
		data := msg{
			Name:"小王子",
			Message:"hello golang",
			Age:18,
		}
		c.JSON(http.StatusOK,data)
	})

	err := r.Run(":9090") //启动 server
	if err!= nil{
		fmt.Println("server failed")
	}
}