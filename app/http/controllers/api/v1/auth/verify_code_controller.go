package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/requests"
	"gohub/pkg/captcha"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"gohub/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

/*
In summary, this package seems to be responsible for handling various aspects of user verification and authentication,
such as generating and displaying image captchas, sending SMS-based and email-based verification codes,
and responding to client requests related to these verification processes.
It's likely a part of a larger web application where users need to be verified
using different methods before gaining access or performing certain actions.
*/

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
	logger.LogIf(err)
	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s, //这个码用于client显示图片验证码
	})
}

// SendUsingPhone 发送手机验证码 ,cliet将图片验证码，手机号发过来，server验证数据，将验证码存储在redis中，然后响应回复手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	// 1. 验证表单
	data := requests.VerifyCodePhoneRequest{}
	//Validate判断client发过来的数据是否能被解析，VerifyCodePhone验证表单是否符合认证的规则，手机和邮箱的验证规则不同
	//1.http请求对象c，2.数据接受对象request，和handler函数VerifyCodePhone
	if ok := requests.Validate(c, &data, requests.VerifyCodePhone); !ok {
		return
	}

	// 2. 发送 SMS
	if ok := verifycode.NewVerifyCode().SendSMS(data.Phone); !ok {
		response.Abort500(c, "发送短信失败~")
	} else {
		response.Success(c)
	}
}

// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {

	// 1. 验证表单
	//这里接受客户端发来的验证码和邮箱
	data := requests.VerifyCodeEmailRequest{}
	//判断client发送来的数据是否有效
	if ok := requests.Validate(c, &data, requests.VerifyCodeEmail); !ok {
		return
	}

	// 2. 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(data.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 验证码失败~")
	} else {
		response.Success(c)
	}
}
