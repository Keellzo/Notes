package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Arvsfonden!1@tcp(127.0.0.1:3306)/go_notes?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database successfully.")
}

func DbMigrate() {
	if DB == nil {
		log.Fatalf("DB migration failed: DB is nil")
		return
	}
	err := DB.AutoMigrate(&Note{})
	if err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}
}
