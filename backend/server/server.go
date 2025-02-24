package main 

import (
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Valid"))
  })

  log.Println("Server running on port 8080")

  if err := http.ListenAndServe(":8080", mux); err != nil {
    panic(err)
  }
}

