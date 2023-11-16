package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
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
