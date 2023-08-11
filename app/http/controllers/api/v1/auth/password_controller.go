// Package auth 处理用户注册、登录、密码重置
package auth

import (
	v1 "gohub/app/http/controllers/api/v1"
	"gohub/app/models/user"
	"gohub/app/requests"
	"gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByPhoneRequest{}
	//这里的request需要传指针，否则的话：json: Unmarshal(non-pointer requests.ResetByEmailRequest)
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	// 2. 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	//1.创建并检验表单
	request := requests.ResetByEmailRequest{}
	if requests.Validate(c, &request, requests.ResetByEmail) == false {
		return
	}

	//2.更新密码
	//获取数据库中用户的信息
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		//成功查询到用户信息之后，将http提交表单中的数据更新到数据库中
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
