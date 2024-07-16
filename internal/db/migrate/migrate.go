package db

import (
	dbModels "github.com/codescalersinternships/secret-note-api-spa-nabil/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string) (*gorm.DB, error) {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = DB.AutoMigrate(&dbModels.User{}, &dbModels.Note{})
	return DB, err
}
