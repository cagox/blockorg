package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/cagox/blockorg/config"
  "github.com/cagox/blockorg/database"
  "github.com/cagox/blockorg/api"
)


func main() {
  database.DialMongoSession()
  defer config.Config.MongoSession.Close()

  api.Routes()

  fmt.Println("I don't actually do anything yet.")


  log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
