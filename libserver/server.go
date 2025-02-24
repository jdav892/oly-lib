package main 

import (
  "net/http"
)

//to modularize middleware
func middleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(write http.ResponseWriter), read *http.Request){
    write.Header().Set("Content-Type", "application/json")
    next.ServeHttp(write, read)
  }
}
