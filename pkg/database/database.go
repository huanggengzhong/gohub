package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
var SQLDB *sql.DB

func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println("数据库链接错误:", err.Error())
	}
	// 底层sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println("数据库底层错误:", err.Error())
	}

	fmt.Println("数据库链接成功")
}
