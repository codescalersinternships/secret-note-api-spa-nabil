package main

import (
	"fmt"
	"log"

	secretnote "github.com/codescalersinternships/secret-note-api-spa-nabil/api/secretnotehandler"
	db "github.com/codescalersinternships/secret-note-api-spa-nabil/internal/db/migrate"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	DB := db.Init()
	r := secretnote.NewServer(DB)
	// Listen and Server in 0.0.0.0:8090
	err = r.Start(":8090")
	if err != nil {
		log.Fatal("server run error:", err)
	}
	fmt.Println("\nGracefully shutting down HTTP server...")
	fmt.Println("Shutdown complete.")
}
