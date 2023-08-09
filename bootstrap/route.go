// bootstrap 包 处理程序初始化逻辑
package bootstrap

import (
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func SetupRoute(router *gin.Engine) {
	//注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册API路由
	routes.RegisterAPIRoutes(router)

	//处理404路由
	setup404Handler(router)
}

// 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		// gin.Logger(),
		middlewares.Logger(), //函数的返回值是 ： gin.HandlerFunc ， 也就是func(*Context)，这个类型的函数就是middleware
		// gin.Recovery(),
		middlewares.Recovery(),
	)

}

// 处理404请求
func setup404Handler(r *gin.Engine) {
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

}
