// Package validators 存放自定义规则和验证器
package validators

import (
	"errors"
	"fmt"
	"github.com/thedevsaddam/govalidator"
	"gohub/pkg/database"
	"strings"
)

// 此方法会在初始化时执行，注册自定义表单验证规则
func init() {
	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,32 排除用户掉 id 为 32 的用户

	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")
		// 第一个参数，表名称，如 users
		tableName := rng[0]
		// 第二个参数，字段名称，如 email 或者 phone
		dbFiled := rng[1]
		// 第三个参数，排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[3]
		}
		requestValue := value.(string)
		query := database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue)
		// 如果传参第三个参数，加上 SQL Where 过滤
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}
		var count int64
		query.Count(&count)
		// 如果找到有 验证不通过，数据库能找到对应的数据
		if count != 0 {
			//看有没有自定义消息
			if message != "" {
				return errors.New(message)
			} else {
				return fmt.Errorf("%v 已经存在,请修改后再重新提交", requestValue)
			}
		}
		return nil
	})
}
