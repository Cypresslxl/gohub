package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

// Phone
type VerifyCodePhoneRequest struct {
	// The jsontag specifies that when this struct is encoded to JSON, the field should be named"captcha_id".
	// The omitemptyoption in thejsontag indicates that if the field is empty, it should be omitted in the JSON output.
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	// The jsontag specifies that the JSON key for this field should be"phone", and the omitemptyoption is used to skip including the field in the JSON if it's empty.
	// The valid` tag could define validation rules specific to phone numbers.
	Phone string `json:"phone,omitempty" valid:"phone"`
}

// VerifyCodePhone 验证表单，返回长度等于零即通过
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {

	// 1. 定制认证规则
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages) //validate 是真正调用第三方库API

	// the provided code snippet data := data.(*VerifyCodePhoneRequest) is a type assertion in Go. It's used for type conversion or type assertion,
	// allowing you to access the underlying value of an interface when you know its concrete type.
	// The code data := data.(*VerifyCodePhoneRequest) is a way to convert the variable data from an interface type to the concrete type VerifyCodePhoneRequest.
	//  This allows you to access the fields and methods specific to the VerifyCodePhoneRequest type on the data variable.
	// 图片验证码
	_data := data.(*VerifyCodePhoneRequest)                                       //assertion and convertion
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs) //比对验证码

	return errs
}

// Email
type VerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`

	Email string `json:"email,omitempty" valid:"email"`
}

// VerifyCodeEmail 验证表单，返回长度等于零即通过
func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {

	// 1. 定制认证规则
	rules := govalidator.MapData{
		"email":          []string{"required", "min:4", "max:30", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"captcha_id": []string{
			"required:图片验证码的 ID 为必填",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	// 图片验证码
	_data := data.(*VerifyCodeEmailRequest) //conversion and assertion
	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
