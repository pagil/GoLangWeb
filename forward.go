package main

import (
  "net/http"
  "io/ioutil"
  "strings"
)

func forward_message(body string) []byte {
  resp, err := http.Post("http://host:80/api/v1/coreteam/notify", "application/json", strings.NewReader(body))
  failOnError(err, "Failed to submit POST Request")

  respBody, err := ioutil.ReadAll(resp.Body)
  failOnError(err, "Failed to read the response")
  return respBody
}
