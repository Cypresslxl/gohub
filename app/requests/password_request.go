package requests

import (
	"gohub/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

// ResetByPhone 验证表单，返回长度等于零即通过
func ResetByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	errs := validate(data, rules, messages)

	// 检查验证码
	_data := data.(*ResetByPhoneRequest) //assertion and conversion
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password" `
}

func ResetByEmail(data interface{}, c *gin.Context) map[string][]string {
	//1.rules
	rules := govalidator.MapData{
		"email":       []string{"required", "min:4", "max:30", "email"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}

	//2.messages
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	//3.validate
	errs := validate(data, rules, messages)

	//数据格式检验成功后，先assertion and conversion，
	_data := data.(*ResetByEmailRequest)
	//再检验邮箱地址和邮箱验证码是否匹配
	validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}
