package user

import (
	"fmt"
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

// 用户模型
type User struct {
	//omitempty字段作用如果值为空省略显示该字段和和该字段内容
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
	res := hash.BcryptCheck(_password, userModel.Password)
	fmt.Println("-----结果,对比密码,新旧是-----", res, _password, userModel.Password)
	return res
}

// 保存用户(比如更新密码)
func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
