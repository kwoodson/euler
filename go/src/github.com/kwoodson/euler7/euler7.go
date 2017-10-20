package main
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number

import (
  "fmt"
  "github.com/kwoodson/mylib"
)

func main() {

  var count int = 1
  //for i := int64(1); i <= 5; i++ {
  for i := int64(1); i <= 1000000; i+=2 {
    if mylib.Isprime(i) {
//fmt.Printf("Prime: %d\n", i)
      count += 1
    }

    if count == 10001 {
        fmt.Printf("Prime: %d\nDone.\n", i)
        break
    }
  }
}
