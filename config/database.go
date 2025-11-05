package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("coffee.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("не удается подключиться к базе данных")
	}
	DB = database
	log.Println("Установлено подключение к базе данных")
}
