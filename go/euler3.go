package main

import (
	"fmt"
	"math"
)

func isprime(n float64) bool {
  for i := 2.0; i < math.Sqrt(n); i++ {
    if (math.Mod(n, i) == 0) {
      return false
    }
  }
  return true
}

func main() {
  var big_num float64 = 600851475143
  var z float64 = 0
  for i := 1.0; i < big_num / 2; i+=2 {
    if (math.Mod(big_num, i) == 0) {
      z = big_num / i
      if (isprime(z)){
        fmt.Printf("Prime: %f\n", z)
        break
      }
    }
  }
}
