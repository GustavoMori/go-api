package main

import (
	"fmt"
	"go-api/db"
	"net/http"
)

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /players", getPlayers)

	mux.HandleFunc("GET /players/{id}", getPlayerID)

	mux.HandleFunc("/", mainRoute)

	db.Connection()

	http.ListenAndServe(":5050", mux)
}
