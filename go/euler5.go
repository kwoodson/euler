package main
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// what is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

import (
	"fmt"
        "os"
)

func isdivisible1_20(n int) bool {

  for i := 20; i > 0; i-- {
    if (n % i != 0) {
      return false
    }
  }
  return true
}

func main() {
  for i := 10000000; i > 0; i++ {
    if isdivisible1_20(i) {
      fmt.Printf("lowest: %d\n", i)
      os.Exit(0)
    }
  }
}
