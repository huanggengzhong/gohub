package category

import (
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"
)

func Get(idStr string) (c Category) {
	database.DB.Where("id", idStr).First(&c)
	return
}

func GetBy(field, value string) (c Category) {
	database.DB.Where("? = ?", field, value).First(&c)
	return
}

func All() (c []Category) {
	database.DB.Find(&c)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Category{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Category{}),
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
