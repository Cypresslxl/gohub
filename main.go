package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default() //Default 返回的是一个 Engine 对象。且默认帮我们注册了两个中间件，Logger 和 Recovery 中间件。
	r := gin.New() //* gin.Engine

	// register middleware
	// r.Use(gin.Logger(), gin.Recovery())
	//程序初始化，初始化路由绑定
	bootstrap.SetupRoute(r)

	err := r.Run(":8848")
	if err != nil {
		//错误处理，端口被占用了或者其他错误

		fmt.Println(err.Error())
	}
	// r.Run() 默认端口为8080
}
