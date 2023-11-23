package topic

import (
	"gohub/app/models"
	"gohub/app/models/category"
	"gohub/app/models/user"
	"gohub/pkg/database"
)

type Topic struct {
	models.BaseModel
	Title      string `json:"title,omitempty" gorm:"type:varchar(255);not null;index"`
	Body       string `json:"body,omitempty" gorm:"type:longtext;not null"`
	UserID     string `json:"user_id,omitempty" gorm:"type:bigint;not null;index"`
	CategoryID string `json:"category_id,omitempty" gorm:"type:bigint;not null;index"`

	//通过 user_id 关联用户
	User user.User `json:"user"`
	// 通过 category_id 关联分类
	Category category.Category `json:"category"`
	models.CommonTimestampsField
}

func (t *Topic) Create() {
	database.DB.Create(&t)
}

func (t *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&t)
	return result.RowsAffected
}

func (t *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&t)
	return result.RowsAffected
}
