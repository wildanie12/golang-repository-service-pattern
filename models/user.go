package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Nama string `json:"nama"`
	Alamat string `json:"alamat"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	StatusUser string
	Gender string
	gorm.Model
}

type Tabler interface {
	TableName() string
}

func (user User) TableName() string {
	return "user"
}