package main

import (
  "fmt"
  "2-packages/neighbor" // [modulename]/[packagename]
)

func main () {
  greet() // this works home.go -> greet() func, because it's apart of the same package!
  fmt.Println() // spacing
  neighbor.Greet()

  /*
   * neighbor.greet() // ./main.go:yy:xx: undefined: neighbor.greet
   * this occurs because greet is not exported, because it's not capitalized
  */
}
