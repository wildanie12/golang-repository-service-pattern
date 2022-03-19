package repositories

import (
	"go-gorm-sandbox/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}


func (repo UserRepository) FindAll() []models.User {
	users := []models.User{}
	repo.DB.Find(&users)
	return users
}

func (repo UserRepository) Find(id uint) models.User {
	user := models.User{}
	repo.DB.Find(&user, id)
	return user
}

func (repo UserRepository) Create(user models.User) models.User {
	repo.DB.Create(&user)
	return user
}

func (repo UserRepository) Update(id uint, user models.User) models.User {
	repo.DB.Save(&user)
	return user
}

func (repo UserRepository) Delete(id uint) int {
	repo.DB.Delete(&models.User{}, id)
	return 1
}