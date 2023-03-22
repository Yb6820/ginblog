package main

import (
	"fmt"
	"ginblog/models"
	"ginblog/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", utils.DbUser, utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)))
	if err != nil {
		fmt.Println("链接数据库失败！请检查参数:\n", err)
	}
	user := models.User{}
	db.Where("username = ?", "zhangsan").First(&user)
	fmt.Println(user)
}
