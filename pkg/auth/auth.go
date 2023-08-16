// Package auth 授权相关逻辑
package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/pkg/logger"
)

// Attempt 尝试登录
func Attempt(loginID string, password string) (user.User, error) {
	userModel := user.GetByMulti(loginID)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) { //不能直接将用户提交密码的和数据库里的密码对比，因为数据库里的密码是被加密过了的
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 登录指定用户
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone) //通过手机号查询数据库获取指定用户信息
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
