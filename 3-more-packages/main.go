package main

import (
  "fmt"
  "3-more-packages/system"
)

func main() {
  fmt.Println("Getting system diagnostics:")
  fmt.Println("Disk space (free/total): " + system.GetDiskSpace() + "MB")
}
