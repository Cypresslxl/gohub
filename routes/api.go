package routes

import (
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			signup := new(auth.SignupController)
			// 判断手机号是否已经注册
			authGroup.POST("/signup/phone/exist", signup.IsPhoneExist)
			// 判断email是否已经注册
			authGroup.POST("signup/email/exist", signup.IsEmailExist)
			//用手机号/邮箱注册账号
			authGroup.POST("/signup/using-phone", signup.SignupUsingPhone)
			authGroup.POST("/signup/using-email", signup.SignupUsingEmail)
			// 发送验证码
			verify := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", verify.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", verify.SendUsingPhone)
			authGroup.POST("verify-codes/email", verify.SendUsingEmail)

		}
		// v1.GET("/", func(c *gin.Context) {
		// JSON 格式相应
		// c.JSON(http.StatusOK, gin.H{
		// "code":    1,
		// "message": "this is v1",
		// })
		// })
	}
}
