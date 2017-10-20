package main

import (
	"fmt"
	//"math"
)

func isprime(n int64) bool {
  if (n <= 2) { return false }
  if (n <= 3) { return true }

  for i := int64(5); i*i < n; i+=6 {
    if (n % i == 0 || n % (i+2) == 0) {
      return false
    }
  }
  return true
}

func main() {
  var big_num int64 = 600851475143
  var z int64 = 0
  for i := int64(3); i < big_num / 2; i+=2 {
    if (big_num % i == 0) {
      z = big_num / i
      if (isprime(z)){
        fmt.Printf("Prime: %d\n", z)
        break
      }
    }
  }
}
