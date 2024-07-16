package db

import (
	"log"
	"os"

	dbModels "github.com/codescalersinternships/secret-note-api-spa-nabil/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't open database")
	}
	DB.AutoMigrate(&dbModels.User{}, &dbModels.Note{})
	return DB
}
