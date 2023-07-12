package main

import (
  "os"
  "io"
  "net/http"
)

func Handler( w http.ResponseWriter, r *http.Request) {
  f, _ := os.Open("./menu")
  io.Copy(w, f)
}

func main() {
  http.HandleFunc("/", Handler)
  http.ListenAndServe(":3000", nil)
}

