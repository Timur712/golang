package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=salam dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// // Добавляем миграцию для модели Task
	// err = DB.AutoMigrate(&taskService.Task{})
	// if err != nil {
	// 	log.Fatalf("Ошибка миграции: %v", err)
	// }
}

