package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default() //Default 返回的是一个 Engine 对象。且默认帮我们注册了两个中间件，Logger 和 Recovery 中间件。
	r := gin.New() //* gin.Engine

	// register middleware
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "sucess",
		})
	})

	// 处理404请求
	r.NoRoute(func(cx *gin.Context) {
		// 获取标头信息的 Accept信息
		head_Accept := cx.Request.Header.Get("Accept")
		if strings.Contains(head_Accept, "text/html") {
			//如果是 HTML 的话
			cx.String(http.StatusNotFound, "404")
		} else {
			cx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})

	r.Run(":8848")
	// r.Run() 默认端口为8080
}
