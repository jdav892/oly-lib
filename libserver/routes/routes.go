package routes

import (
  "libserver/handlers"
  "github.com/gorilla/mux"
)

//RegisterRoutes to set up API routes
func RegisterRoutes() *mux.Router {
  response := mux.NewRouter()

  //defining the routes with path params
  response.HandleFunc("/lib/exercises", handlers.AllExercisesHandler).Methods("Get")
  response.HandleFunc("/lib/exercises/{sport}", handlers.ExercisesByCategoryHandler).Methods("GET")

  return response
}
