package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gohub/app/models/user"
	"gohub/pkg/logger"
)

func AuthLoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("用户不存在")
	}
	return userModel, nil
}

// 手机号/邮箱/用户名+密码尝试登录
func Attempt(mulValue, password string) (user.User, error) {
	userModel := user.GetByMulti(mulValue)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}

// 从gin.context中获取用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("context无法获取用户"))
		return user.User{}
	}
	return userModel
}

// 从gin.context中获取用户ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
