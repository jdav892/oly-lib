package handlers 

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

//predefined exercise categories as template for json repsonses, will seed pg db later
var exerciseCategories = map[string][]string{
  "powerlifting": {"Bench Press", "Deadlift", "Back Squat"},
  "weightlifting": {"Snatch", "Clean and Jerk"},
  "general-fitness": {"Front Squat", "Strict Press", "Barbell Row"},
  "plyometrics": {"Box Jump", "Broad Jump", "Hurdles"},
}

//retrieves all exercises and their categories
func AllExercisesHandler(write http.ResponseWriter, read *http.Request) {
  //responds with all exercises
  response := map[string]interface{}{
    "exercises": exerciseCategories,
  }
  //encodes and returns as JSON 
  write.Header().Set("Content-Type", "application/json")
  json.NewEncoder(write).Encode(response)
}


//retrieves exerciese by categories
func ExercisesByCategoryHandler(write http.ResponseWriter, read *http.Request) {
  //get the sport from url path(captured as path parameter)
  vars := mux.Vars(read)
  sport := vars["sport"]

  //fetch exercises from specified sport
  exercises, exists := exerciseCategories[sport]
  if !exists {
    http.Error(write, "Sport not found", http.StatusNotFound)
    return
  }

  //return exercise as JSON
  response := map[string]interface{}{
    "sport": sport,
    "exercises": exercises,
  }
  write.Header().Set("Content-Type", "application/json")
  json.NewEncoder(write).Encode(response)
}


