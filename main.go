package main

import (
  "net/http"
  "bytes"
  "io"
)

func main() {
  // ServerMux
  http.HandleFunc("/health", health)
  http.HandleFunc("/send", send)
  http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`{ "Status": "Up"}`))
}

func send(w http.ResponseWriter, r *http.Request) {
  //decoder := json.NewDecoder(req.Body)
  bodyString := streamToString(r.Body)
  send_message(bodyString)
  response := forward_message(bodyString)
  w.Write(response)
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
