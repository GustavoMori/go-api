package main

import (
	"fmt"
	"net/http"
)

func getPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET of players")
	w.Write([]byte("Welcome to Go as webserver"))
}

func getPlayerID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Println("GET player with id = " + id)
	w.Write([]byte("Player with id = " + id))
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main route")
	w.Write([]byte("Main route"))
}
