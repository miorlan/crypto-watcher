package main

import (
	"cryptoWatcher/internal/api"
	"cryptoWatcher/internal/config"
	"cryptoWatcher/internal/database"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitDB(cfg)
	defer db.Close()
	r := api.SetupRouter(db)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
