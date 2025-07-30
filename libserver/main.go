package main 

import (
  "fmt"
  "log"
  "net/http"
	
	"github.com/rs/cors"
	"libserver/middleware"
  "libserver/routes"
)

func main() {

	port := ":8080"
	router := routes.RegisterRoutes()
	handler := middleware.JSONMiddleware(router)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(handler)

  fmt.Println("Server running on localhost" + port)
  log.Fatal(http.ListenAndServe(port, corsHandler))
}
