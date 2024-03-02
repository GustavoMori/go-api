package main

import (
	"fmt"
	"go-api/db"
	"go-api/internal/handlers"
	"net/http"
)

func init() {
	fmt.Println("init")
}

func main() {
	mux := http.NewServeMux()

	playerRoutes(mux)

	mux.HandleFunc("/", handlers.MainRoute)

	dsn := db.MakeDSN("localhost", "myuser", "mydatabase", "mypassword")
	db.InitDB(dsn)
	db.RunMigrates()

	http.ListenAndServe(":5050", mux)
}

func playerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /players", handlers.GetPlayers)

	mux.HandleFunc("GET /players/{id}", handlers.GetPlayerByID)

	mux.HandleFunc("POST /players/create", handlers.CreatePlayer)

	mux.HandleFunc("PUT /players/{id}", handlers.UpdatePlayer)

	mux.HandleFunc("DELETE /players/{id}", handlers.DeletePlayer)
}
