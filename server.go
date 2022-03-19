package main

import "gorm.io/gorm"

// "database/sql"
// "fmt"

// _ "github.com/go-sql-driver/mysql"

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	TanggalLahir string `json:"tanggal_lahir"`
	StatusUser string
	Gender string
	CreatedAt string
	UpdatedAt string
}

func InitialMigration() {
	DB.AutoMigrate(User)
}




// func findAll() []User {
// 	db := connect()

// 	result, err := db.Query("SELECT * FROM user")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	dataset := []User{}
// 	for result.Next() {
// 		user := User{}
// 		result.Scan(&user.id, &user.nama, &user.alamat, &user.tanggal_lahir, &user.status_user, &user.gender, &user.created_at, &user.updated_at)
// 		dataset = append(dataset, user)
// 	}
// 	return dataset
// }


// func create(user User) User {
// 	db := connect()

// 	query := fmt.Sprintf("INSERT INTO user (nama, alamat, tanggal_lahir, status_user, gender, created_at, updated_at) VALUES (``)")
// 	result, err := db.Exec()
// }


// func main() {
// 	dataUser := findAll()
// 	fmt.Println(dataUser)
// }
