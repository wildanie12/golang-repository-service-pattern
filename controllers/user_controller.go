package controllers

import (
	"fmt"
	"go-gorm-sandbox/config"
	"go-gorm-sandbox/repositories"
	"go-gorm-sandbox/services"
	"go-gorm-sandbox/utilities"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = config.MysqlConnectGorm()
}

func UserIndex() {
	repo := repositories.UserRepository{
		DB: db,
	}
	data := services.UserAll(repo)
	fmt.Println(utilities.JsonEncode(data))
}