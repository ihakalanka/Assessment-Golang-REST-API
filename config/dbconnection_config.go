package config

import (
	"os"
	"github.com/ihakalanka/Assessment-Golang-REST-API/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/ihakalanka/Assessment-Golang-REST-API/utils"
)

var DB *gorm.DB

func ConnectDB() {
	InitEnv()
	dsn := os.Getenv("DSN")
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}