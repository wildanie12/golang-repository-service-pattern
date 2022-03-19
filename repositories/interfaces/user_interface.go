package interfaces

import "go-gorm-sandbox/models"

type UserInterface interface {
	FindAll() []models.User 
	Find(id uint) models.User
	Create(user models.User) models.User
	Update(id uint, user models.User) models.User
	Delete(id uint) int 
}