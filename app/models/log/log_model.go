package log

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Log struct {
	models.BaseModel
	Message string `json:"message,omitempty" gorm:"type:longtext;not null"`
	models.CommonTimestampsField
}

func (t *Log) Create() {
	database.DB.Create(&t)
}

func (t *Log) Save() (rowsAffected int64) {
	result := database.DB.Save(&t)
	return result.RowsAffected
}

func (t *Log) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&t)
	return result.RowsAffected
}
