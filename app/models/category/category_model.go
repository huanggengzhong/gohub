package category

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Category struct {
	models.BaseModel
	Name        string `json:"name,omitempty" gorm:"type:varchar(255);not null;index"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255);default:null"`
	models.CommonTimestampsField
}

func (c *Category) Create() {
	database.DB.Create(&c)
}

func (c *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&c)
	return result.RowsAffected
}

func (c *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&c)
	return result.RowsAffected
}
