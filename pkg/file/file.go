package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gohub/pkg/app"
	"gohub/pkg/auth"
	"gohub/pkg/helpers"
	"mime/multipart"
	"os"
	"path/filepath"
)

// 保存头像 // public/uploads/avatars/2021-12-22/1/nFDacgaWKpWWOmOt.png
func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {

	// 确保目录存在，不存在创建
	dirPath := fmt.Sprintf("public/uploads/%s/%s/", app.TimenowInTimezone().Format("2006-01-02"), auth.CurrentUID(c))
	os.MkdirAll(dirPath, 0755) //0755：所有者有读、写、执行权限，其他人有读、执行权限。

	//保存文件 (在Gin框架中，c.SaveUploadedFile方法不是直接由Gin提供的，而是由标准库的net/http和mime/multipart包提供的。Gin建立在这些标准库之上，因此你可以使用标准库的相关功能来处理文件上传)
	fileName := randomFileName(file)
	avatarPath := dirPath + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return "", err
	}
	return avatarPath, nil

}

// 随机生成文件名
func randomFileName(file *multipart.FileHeader) string {
	return helpers.RandomNumber(16) + filepath.Ext(file.Filename) //Ext方法获取扩展名
}
