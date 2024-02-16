package main

import (
	"fmt"
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

	http.ListenAndServe(":5050", mux)
}
