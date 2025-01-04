package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println("Getting secret key...")
  keyFileName := "secret-key.txt"
  content, err := os.ReadFile(keyFileName)
  if err != nil {
    fmt.Println("Secret key: Failed to retrieve")
  }
  fmt.Println("Secret key: " + string(content))
}
