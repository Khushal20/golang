package repository

import (
	"fmt"
	"khushal/goMux/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dbConfig := "host=localhost user=postgres password=root dbname=user port=5432"
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error while connecting to DB %v", err.Error())
	}
	db.AutoMigrate(&entity.User{})
	return db, nil
}
