package repository

import (
	"khushal/goMux/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func InitDB() *gorm.DB {
	dbConfig := "host=localhost user=postgres password=root dbname=user port=5432"
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		panic("Error while connecting to DB")
	}
	db.AutoMigrate(&entity.User{})
	return db
}
