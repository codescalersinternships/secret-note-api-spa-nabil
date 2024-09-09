package db

import (
	"fmt"

	db "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/internal/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dsn string) (db.Store, error) {
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = DB.AutoMigrate(&db.User{}, &db.Note{})
	temp := &db.SqlStore{}
	temp.GormStore = DB
	if err == nil {
		fmt.Println("\nServer migration is done now...")
	}
	return temp, err
}
