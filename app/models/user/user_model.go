package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

// 用户模型
type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
}

// Create创建用户
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// 对比密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
