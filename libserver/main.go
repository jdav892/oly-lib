package main

import (
	"fmt"
	"log"
	"net/http"

	"libserver/middleware"
	"libserver/routes"

	"github.com/rs/cors"
)

func main() {

	port := ":8080"
	router := routes.RegisterRoutes()
	handler := middleware.JSONMiddleware(router)

	//probably going to add more to this
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
	}).Handler(handler)

	fmt.Println("Server running on localhost" + port)
	log.Fatal(http.ListenAndServe(port, corsHandler))
}
