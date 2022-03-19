package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlConnectGorm() *gorm.DB {
	dsn := "root:root@tcp(localhost:3306)/exp_go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}