package main

import "net/http"

func main() {
  // ServerMux
  http.HandleFunc("/", hello)
  http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<html>"))
  w.Write([]byte("<body>"))
  w.Write([]byte("<h1>Hello!</h1>"))
  w.Write([]byte("</body>"))
  w.Write([]byte("</html>\r\n"))
}
