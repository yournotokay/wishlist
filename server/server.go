package server

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

// Starts our server on the given port
func StartMyApp(port int) {
  http.HandleFunc("/", handler)
  http.ListenAndServe(fmt.Sprintf(":%v",port), nil)
}
