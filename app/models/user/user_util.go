package user

import "gohub/pkg/database"

// 根据email判断用户是否已经存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// 根据phone判断用户是否已经存在

func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// 根据手机号返回用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// 通过手机号/Email/用户名 返回用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.Where("phone = ?", loginID).Or("email = ?", loginID).Or("name = ?", loginID).First(&userModel)
	return
}

// 通过ID  返回用户
func Get(id string) (userModel User) {
	database.DB.Where("id", id).First(&userModel)
	return
}

//所有用户
func All() (users []User) {
	database.DB.Find(&users)
	return
}