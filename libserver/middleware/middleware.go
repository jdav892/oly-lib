package middleware

import (
  "net/http"
)

//to modularize middleware
func JSONMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    next.ServeHTTP(w, r)
  })
}
