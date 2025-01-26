package server 
import (
  "encoding/json"
  "log"
  "net/http"
)
type Todo struct {
  ID  string `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}
var todos []Todo
func main() {
  // Define routes
  http.HandleFun("/todos", getTodos)
  //Start the server
  log.Fatal(http.ListenAndServer(":8080"), nil)
}
func getTodos(w http.ResponseWritter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(todos)
}
