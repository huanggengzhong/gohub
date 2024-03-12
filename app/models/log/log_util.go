package log

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func Get(idstr string) (log Log) {
	database.DB.Where("id", idstr).First(&log)
	return
}

func GetBy(field, value string) (log Log) {
	database.DB.Preload(clause.Associations).Where("? = ?", field, value).First(&log)
	return
}

func All() (logs []Log) {
	database.DB.Find(&logs)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Log{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (logs []Log, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Log{}),
		&logs,
		app.V1URL(database.TableName(&Log{})),
		perPage,
	)
	return
}
