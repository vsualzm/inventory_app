package main

import (
	"time"

	"gorm.io/gorm"
)

// untuk model

type Invetory struct {
	ID        int `gorm:"primary_key"`
	NamaItem  string
	Quantity  int
	Category  string
	Price     float64
	CreatedAt time.Time
	UpdateAt  time.Time
}

var DB *gorm.DB

// func DatabaseConnection() {
// 	// koneksi ke database

// 	dsn := "host=localhost user=postgres password=1234 dbname=root port=5432 sslmode=disable"
// 	db, err := gorm.Open("postgres", dsn)

// 	if err != nil {
// 		panic("Failed to connect to database")
// 	}

// 	db.AutoMigrate(&Invetory{})
// 	DB = db

// }

// func main() {

// 	// bermain dengan crud dengan gorm

// }
