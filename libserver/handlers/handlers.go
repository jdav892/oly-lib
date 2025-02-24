package handlers 

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

//predefined exercise categories
var exerciseCategories = map[string][]string{
  "powerlifting": {"Bench Press", "Deadlift", "Back Squat"},
  "weightlifting": {"Snatch", "Clean and Jerk"},
  "strongman": {"Atlas Stone", "Axle Clean and Press", "Yoke Carry"},
  "general exercies": {"Front Squat", "Strict Press", "Barbell Row"},
}


//retrieves exerciese by categories
func ExercisesByCategoryHandler(w http.ResponseWriter, r *http.Request) {
  //get the sport from url path(captured as path parameter)
  vars := mux.Vars(r)
  sport := vars["sport"]

  //fetch exercies from specified sport
  exercises, exists := exerciseCategories[sport]
  if !exists {
    http.Error(w, "Sport not found", http.StatusNotFound)
    return
  }

  //return exercise as JSON
  response := map[string]interface{}{
    "sport": sport,
    "exercises": exercises,
  }

  json.NewEncoder(w).Encode(response)
}


