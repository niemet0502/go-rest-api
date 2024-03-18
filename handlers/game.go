package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-server/entities"
	"go-server/services"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Write response for the home page
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Marius The last codebender"))
}

func HandlerGetGames(w http.ResponseWriter, r *http.Request) {
	log.Println("Call the /games [GET] endpoint")
	gameService := services.CreateNewGameService()

	games := gameService.FindAll()

	jsonResponse, _ := json.Marshal(games)

	// Set Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response body
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Call the /games/{id} endpoint")
	vars := mux.Vars(r)

	id := vars["id"]

	num, _ := strconv.Atoi(id)

	log.Println(num)

	gameService := services.CreateNewGameService()

	game, err := gameService.GetById(num)

	if err != nil {
		log.Println("Game not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonResponse, _ := json.Marshal(game)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func HandleCreateGame(w http.ResponseWriter, r *http.Request) {
	var game entities.Game
	log.Println("Call the /games [POST] endpoint")

	err := json.NewDecoder(r.Body).Decode(&game)

	if err != nil {
		log.Fatal("Failed to decode the payload")
	}

	gameService := services.CreateNewGameService()

	err = gameService.Create(&game)

	jsonResponse, _ := json.Marshal(game)

	if err != nil {
		http.Error(w, "Failed to create game", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
