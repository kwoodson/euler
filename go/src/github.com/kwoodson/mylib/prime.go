package mylib

import (
  "math"
)

func Isprime(n int64) bool {
  if (n < 2 || n % 2 == 0) { return false }
  if (n == 2) { return true }
  if (n == 3) { return true }
  if (n % 3 == 0) { return false }

  var stop int64 = int64(math.Pow(float64(n), .5))

  for i := int64(5); i <= stop; i+=6 {
    if (n % i == 0 || n % (i+2) == 0) {
      return false
    }
  }
  return true
}
