package app

import (
	"fyque/model/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres dbname=fyque port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.User{}, &domain.Book{})

	return db, nil
}