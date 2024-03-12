package bootstrap

import (
	"fmt"
	"gohub/app/models/category"
	"gohub/app/models/log"
	"gohub/app/models/topic"
	"gohub/app/models/user"
	"gohub/pkg/config"
	"gohub/pkg/database"
	"gohub/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDB() {
	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		//构建DSN信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic("缺少数据库")

	}
	//数据库链接 并设置gorm日志格式
	database.Connect(dbConfig, logger.NewGormLogger())
	//database.Connect(dbConfig, logger.Default.LogMode(logger.Info))
	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)
	//根据用户模型自动创建数据库表
	database.DB.AutoMigrate(&user.User{})
	//根据分类模型自动创建数据库表
	database.DB.AutoMigrate(&category.Category{})
	//根据话题模型自动创建数据库表
	database.DB.AutoMigrate(&topic.Topic{})
	//根据日志模型自动创建数据库表
	database.DB.AutoMigrate(&log.Log{})
}
