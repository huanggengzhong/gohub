package models

import "gohub/app/models"

// 用户模型
type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
}
