package main 

import (
  "fmt"
  "log"
  "libserver/routes"
  "net/http"
)

func main() {

  port := ":8080"
  mux := routes.RegisterRoutes()

  fmt.Println("Server running on localhost:" + port)
  log.Fatal(http.ListenAndServe(port, mux))
}
