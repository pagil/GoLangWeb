package main

import (
  "net/http"
  "bytes"
  "io"
  "encoding/json"
  "log"
)

type user_struct struct {
  FirstName string
  LastName string
  Email string
  Phone string
}

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

  var u user_struct
  err := json.Unmarshal([]byte(bodyString), &u)
  failOnError(err, "Failed to unmarshal message from the queue")

  postString := `
  {
    "ToEmail": "` + u.Email + `",
    "FromEmail": "go.lang.demo.cs@gmail.com",
    "FromDisplayName": "PIT STOP",
    "ClientId": "CoreTeam",
    "ApplicationId": "Survey",
    "TemplateName": "CaesarSellerad",
    "DeliveryTemplates": [
      {
        "DeliveryMode": "Email",
        "Data": [
          {
            "Name": "email",
            "Properties": [
              {
                "Name": "contenthost",
                "Value": "http://surveymuffin.web.csnb532.csdev.com.au"
              },
              {
                "Name": "name",
                "Value": "` + u.FirstName + `"
              }
            ],
            "Data": [
              "TemplateData"
            ]
          }
        ]
      }
    ]
  }`
  log.Printf("> Request: POSTing string to EMAIL API: %s", postString)
  response := forward_message(postString)
  log.Printf("> Response: %s", response)
  // allow cross domain AJAX requests
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(response)
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
