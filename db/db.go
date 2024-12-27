package db

import (
	"go-mysql-phpmyadmin/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// MySQL connection string: username:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:11111111@tcp(127.0.0.1:3306)/gorm_db?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Automigrate models
	err = DB.AutoMigrate(&models.Event{}, &models.Registration{})
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}
}
