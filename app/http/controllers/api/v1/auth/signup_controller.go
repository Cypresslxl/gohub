// 处理用户身份认证相关逻辑
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (is *SignupController) IsPhoneExist(cx *gin.Context) {
	// panic("this is panic test")
	// 初始化请求对象
	request := requests.SignupPhoneExistRequest{}

	// validate 判断数据格式是否合法
	// cx代表的是http请求，request代表的是处理后的数据，第三个参数代表的是处理函数handler
	if ok := requests.Validate(cx, &request, requests.ValidateSignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	cx.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// IsEmailExist 检测邮箱是否已注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
		return
	}

	//  检查数据库并返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
