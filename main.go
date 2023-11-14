package main

import (
 //  "errors"
	"fmt"
	"net/http"
	"os"
  "database/sql"

  _ "github.com/lib/pq"
)
func main() {

  port := os.Getenv("PORT")
  if port == "" {
    port = ":8080"
  } else {
    port = ":" + port
  }

  http.HandleFunc("/", Redirect)
  err := http.ListenAndServe(port, nil)
  if err != nil {
    fmt.Printf("Error: %s\n", err)
    os.Exit(1)
  }
}
