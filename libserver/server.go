package main 

import (
  "net/http"
)

//to modularize middleware
func middleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter), r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    next.ServeHttp(w, r)
  }
}
