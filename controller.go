package main

import (
	"encoding/json"
	"fmt"
	"go-api/db"
	"io"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type PlayerRequestBody struct {
	Name string `json:"name"`
}

type Player struct {
	gorm.Model
	ID   int
	Name string
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET of players")

	var players []Player
	gorm := db.GormConnection()

	tx := gorm.Find(&players)
	if tx.Error != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	for _, player := range players {
		w.Write([]byte("Player name = " + player.Name))
	}
}

func getPlayerByID(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "NaN ID", http.StatusBadRequest)
		return
	}

	var player = Player{ID: id}
	gorm := db.GormConnection()

	tx := gorm.First(&player)
	if tx.Error != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	fmt.Println("GET player with id = " + strID)
	w.Write([]byte("Player name = " + player.Name))
}

func createPlayer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	var requestBody PlayerRequestBody

	if err := json.Unmarshal(body, &requestBody); err != nil {
		http.Error(w, "Error decoding the request body", http.StatusBadRequest)
		return
	}

	if requestBody.Name == "" {
		http.Error(w, "Name is invalid", http.StatusBadRequest)
		return
	}

	gorm := db.GormConnection()

	gorm.Create(&Player{Name: requestBody.Name})

	fmt.Println("Player with name = " + requestBody.Name + " was created")
	w.Write([]byte("Player with name = " + requestBody.Name + " was created"))
}

func updatePlayer(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "NaN ID", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Error reading the request body", http.StatusBadRequest)
		return
	}

	var requestBody PlayerRequestBody
	if err := json.Unmarshal(body, &requestBody); err != nil {
		http.Error(w, "Error decoding the request body", http.StatusBadRequest)
		return
	}

	if requestBody.Name == "" {
		http.Error(w, "Name is invalid", http.StatusBadRequest)
		return
	}

	var player = Player{ID: id}

	gorm := db.GormConnection()

	tx := gorm.Model(&player).Update("name", requestBody.Name)
	if tx.Error != nil {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	fmt.Println("Updated player with id = " + strID)
	w.Write([]byte("New player name = " + player.Name))
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	strID := r.PathValue("id")

	id, err := strconv.Atoi(strID)
	if err != nil {
		http.Error(w, "NaN ID", http.StatusBadRequest)
		return
	}

	var player = Player{ID: id}
	gorm := db.GormConnection()

	tx := gorm.Scopes(db.NotBeRonaldinho).Model(&player).Update("deleted_at", time.Now())

	fmt.Println(tx.RowsAffected == 0)

	if tx.RowsAffected == 0 {
		http.Error(w, "400", http.StatusBadRequest)
		return
	}

	fmt.Println("Player was deleted")
	w.Write([]byte("Player was deleted"))
}

func mainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Main route")
	w.Write([]byte("Main route"))
}
