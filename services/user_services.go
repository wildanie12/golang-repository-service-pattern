package services

import (
	"go-gorm-sandbox/models"
	"go-gorm-sandbox/repositories/interfaces"
	"time"
)

func UserAll(repo interfaces.UserInterface) []models.User {
	users := []models.User{}
	users = repo.FindAll()
	return users
}

func UserCreate(repo interfaces.UserInterface) {
	user1 := models.User{
		Nama: "Budi",
		Alamat: "Jl. Jakarta",
		TanggalLahir: time.Date(2000, 7, 28, 12, 0, 0, 0, time.FixedZone("Asia/Jakarta", 0)),
		StatusUser: "aktif",
		Gender: "L",
	}
	repo.Create(user1)
}