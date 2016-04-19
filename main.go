package main

import (
  "net/http"
  "bytes"
  "io"
)

func main() {
  // ServerMux
  http.HandleFunc("/", hello)
  http.HandleFunc("/send", send)
  http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<html>"))
  w.Write([]byte("<body>"))
  w.Write([]byte("<h1>Hello!</h1>"))
  w.Write([]byte("</body>"))
  w.Write([]byte("</html>\r\n"))
}

func send(w http.ResponseWriter, r *http.Request) {
  //decoder := json.NewDecoder(req.Body)
  send_message(streamToString(r.Body))
  w.Write([]byte("{\"message\":\"Sent Succesfully!\"}"))
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
