package routes

import (
	"github.com/gorilla/mux"
	"go-server/handlers"
)

func InitRoutes() *mux.Router {
	// Create a new router
	r := mux.NewRouter()

	// Define route handlers
	r.HandleFunc("/", handlers.HomeHandler)

	gamesRouter := r.PathPrefix("/games").Subrouter()
	gamesRouter.HandleFunc("", handlers.HandleCreateGame).Methods("POST")
	gamesRouter.HandleFunc("", handlers.HandlerGetGames).Methods("GET")
	gamesRouter.HandleFunc("/{id}", handlers.HandleGetGame).Methods("GET")
	gamesRouter.HandleFunc("/{id}", handlers.HandleDeleteGame).Methods("DELETE")

	return r
}
