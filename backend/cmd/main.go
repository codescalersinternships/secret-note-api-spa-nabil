package main

import (
	"fmt"
	"log"
	"os"

	secretnote "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/api"
	db "github.com/codescalersinternships/secret-note-api-spa-nabil/backend/internal/db/migrate"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	DB, err := db.Init(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("can't migrate correctly", err)
	}
	r := secretnote.NewServer(DB)
	// Listen and Server in 0.0.0.0:8090
	err = r.Start(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("server run error:", err)
	}
	fmt.Println("\nGracefully shutting down HTTP server...")
	fmt.Println("Shutdown complete.")
}
