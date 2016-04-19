package main

import "net/http"

func main() {
  // ServerMux
  http.HandleFunc("/", hello)
  http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("hello!"))
}
