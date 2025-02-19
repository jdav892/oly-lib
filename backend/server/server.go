package server 
import (
  "encoding/json"
  "log"
  "net/http"
)
type Lift struct {
  ID  string `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}
var lifts []Lift
func main() {
  // Define routes
  http.HandleFun("/lifts", getLifts)
  //Start the server
  log.Fatal(http.ListenAndServer(":8080"), nil)
}
func getTodos(w http.ResponseWritter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(todos)
}
