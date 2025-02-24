package routes

import (
  "libserver/handlers"
  "github.com/gorilla/mux"
)

//RegisterRoutes to set up API routes
func RegisterRoutes() *mux.Router {
  r := mux.NewRouter()

  //defining the routes with path params
  r.HandleFunc("/lib/exercises", handlers.AllExercisesHandler).Methods("Get")
  r.HandleFunc("/lib/exercises/{sport}", handlers.ExercisesByCategoryHandler).Methods("GET")

  return r
}
