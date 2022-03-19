package main

import (
	"go-gorm-sandbox/config"
	"go-gorm-sandbox/controllers"
	"go-gorm-sandbox/models"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	db = config.MysqlConnectGorm()
	db.AutoMigrate(&models.User{})
}


func main() {
	controllers.UserIndex()
}



// func main() {
	
// 	// Get All Data
// 	users := []models.User{}
// 	db.Find(&users)
// 	fmt.Println(utilities.JsonEncode(users))
	
	
// 	// // Get First data
// 	// user := models.User{}
// 	// db.First(&user)
// 	// fmt.Println(user)


// 	// // Create 
// 	// user1 := models.User{
// 	// 	Nama: "Budi",
// 	// 	Alamat: "Jl. Jakarta",
// 	// 	TanggalLahir: "2000/07/28",
// 	// 	StatusUser: "aktif",
// 	// 	Gender: "L",
// 	// }
// 	// result := db.Create(&user1)
// 	// fmt.Println(result)

// 	// ---------------------
// 	// Update
// 	// ---------------------

// 	// Update 
// 	user := models.User{}
// 	db.First(&user, 4)
// 	user.Nama = "Cahyo"
// 	user.Alamat = "Jl. Surabaya"
// 	user.TanggalLahir = "2000/05/28"
// 	user.Gender = "L"

// 	// user.Nama = "Dinda"
// 	// user.Alamat = "Pasuruan"
// 	// user.TanggalLahir = "1999/03/01"
// 	// user.Gender = "P"

// 	db.Save(&user)
// 	fmt.Println(utilities.JsonEncode(user))
// }