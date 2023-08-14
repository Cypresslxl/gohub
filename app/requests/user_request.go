package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/auth"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {

	//查询用户名重复时，过滤掉当前用户ID
	currentID := auth.CurrentUID(c)
	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + currentID},
		"introduction": []string{"min_cn:4", "max_cn:240"},
		"city":         []string{"min_cn:2", "max_cn:20"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误,只允许数字和英文",
			"between：用户名长度需在3~20之间",
			"no_exists:用户名已被占用",
		},
		"introduction": []string{
			"min_cn:描述长度需至少4个字",
			"max_cn:描述长度不超过240个字",
		},
		"city": []string{
			"min_cn:城市需要至少2个字",
			"max_cn:城市不能超过20个字",
		},
	}
	return validate(data, rules, messages)
}
