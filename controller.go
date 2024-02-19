package main

import (
	"encoding/json"
	"fmt"
	"go-api/db"
	"io"
	"net/http"

	"gorm.io/gorm"
)

type CreatePlayerRequestBody struct {
	Name string `json:"name"`
}

type Player struct {
	gorm.Model
	Name string
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET of players")
	w.Write([]byte("Welcome to Go as webserver"))
}

func getPlayerID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	fmt.Println("GET player with id = " + id)
	w.Write([]byte("Player with id = " + id))
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	var requestBody CreatePlayerRequestBody

	if err := json.Unmarshal(body, &requestBody); err != nil {
		http.Error(w, "Error decoding the request body", http.StatusBadRequest)
		return
	}

	if requestBody.Name == "" {
		http.Error(w, "Name is invalid", http.StatusBadRequest)
		return
	}

	entity := db.GormConnection()

	entity.Create(&Player{Name: requestBody.Name})

	fmt.Println("Player with name = " + requestBody.Name + " was created")
	w.Write([]byte("Player with name = " + requestBody.Name + " was created"))
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main route")
	w.Write([]byte("Main route"))
}
