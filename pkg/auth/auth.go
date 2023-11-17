package auth

import (
	"errors"
	"gohub/app/models/user"
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
