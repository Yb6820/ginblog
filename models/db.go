package models

import (
	"fmt"
	"ginblog/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库的函数
var db *gorm.DB
var err error

func InitDb() {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)))
	if err != nil {
		fmt.Println("链接数据库失败！请检查参数:\n", err)
	}
	db.AutoMigrate()
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
