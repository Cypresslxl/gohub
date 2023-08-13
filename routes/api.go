package routes

import (
	controllers "gohub/app/http/controllers/api/v1"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

// 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
		// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
		// 测试时，可以调高一点。
		v1.Use(middlewares.LimitIP("200-H")) //每小时最多200个请求
		{
			authGroup := v1.Group("/auth")
			// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
			// 测试时，可以调高一点
			authGroup.Use(middlewares.LimitIP("1000-H"))
			{
				//1.signup
				signupController := new(auth.SignupController)
				// 判断手机号是否已经注册
				authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), signupController.IsPhoneExist)
				// 判断email是否已经注册
				authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), signupController.IsEmailExist)
				//用手机号/邮箱注册账号
				authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), signupController.SignupUsingPhone)
				authGroup.POST("/signup/using-email", middlewares.GuestJWT(), signupController.SignupUsingEmail)

				// 2.verify,verify-code需要限流，邮箱和手机短信限流更严苛，图片验证码限流相对宽松
				verifyCodeController := new(auth.VerifyCodeController)
				// 图片验证码
				authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("50-H"), verifyCodeController.ShowCaptcha)
				//手机短信和邮箱验证码
				authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), verifyCodeController.SendUsingPhone)
				authGroup.POST("verify-codes/email", middlewares.LimitPerRoute("20-H"), verifyCodeController.SendUsingEmail)

				//3.login
				loginController := new(auth.LoginController)
				// 使用手机号，短信验证码进行登录
				authGroup.POST("/login/using-phone", middlewares.GuestJWT(), loginController.LoginByPhone)
				//使用密码登陆，LoginID可以是phone,email,userName
				authGroup.POST("/login/using-password", middlewares.GuestJWT(), loginController.LoginByPassword)
				//
				//3.2refresh token
				authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), loginController.RefreshToken)

				//4.reset password
				passwordController := new(auth.PasswordController)
				//通过手机号和验证码reset密码
				authGroup.POST("/password-reset/using-phone", middlewares.AuthJWT(), passwordController.ResetByPhone)
				//通过邮箱和验证码reset密码
				authGroup.POST("/password-reset/using-email", middlewares.AuthJWT(), passwordController.ResetByEmail)

				//	5.User
				user := new(controllers.UsersController)
				//h获取当前用户需要Token认证，所以使用AuthJWT()返回的中间件
				authGroup.GET("/user", middlewares.AuthJWT(), user.CurrentUser)
				usersGroup := authGroup.Group("/users")
				{
					usersGroup.GET("", user.Index)
				}
			}

			//category
			category := new(controllers.CategoriesController)
			categoryGroup := v1.Group("/categories")
			{
				categoryGroup.POST("", middlewares.AuthJWT(), category.Store) //登录用户才能创建分类，所以我们用了 AuthJWT 中间件。
				categoryGroup.PUT("/:id", middlewares.AuthJWT(), category.Update)
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
}
