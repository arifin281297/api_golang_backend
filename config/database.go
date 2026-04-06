package config

import (
	"fmt"
	"go-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // <- ini variabel global DB

func ConnectDatabase() {
	dsn := "root:macanbetina21@tcp(127.0.0.1:3306)/crudapigolang?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}

	DB = database

	// AutoMigrate untuk tabel cabang
	if err := DB.AutoMigrate(&models.User{}, &models.UserCabang{}, &models.UserGroup{}); err != nil {
		fmt.Println("Migration failed:", err)
		return
	}

	fmt.Println("Database connected & migrated!")
}
