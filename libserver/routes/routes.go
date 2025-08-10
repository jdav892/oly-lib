package routes

import (
	"libserver/handlers"

	"github.com/gorilla/mux"
)

// RegisterRoutes to set up API routes
func RegisterRoutes() *mux.Router {
	response := mux.NewRouter().StrictSlash(true)

	response.HandleFunc("/lib/exercises", handlers.AllExercisesHandler).Methods("GET", "OPTIONS")

	return response
}
