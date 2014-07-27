package main

import "fmt"

func main() {
  var total int = 0
  for i := 0; i < 1000; i++ {
    if (i % 3 == 0) || (i % 5 == 0) {
      total += i
    }
  }

  fmt.Printf("Total: %d\n", total)
}
