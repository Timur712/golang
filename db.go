package main

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)
// переменная, через которую мы будем работать с БД
var DB *gorm.DB

func InitDB() {

    dsn := "host=localhost user=postgres password=salam dbname=postgres port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Ошибка подключения к БД:%v", err)
    }
    DB.AutoMigrate(&Message{})
}
