package main

import (
  "fmt"
)

func fib(a int) int {
  if (a <  1) {
    return 0
  }
  if (a ==  1) {
    return 1
  }

  return fib(a - 1) + fib(a - 2)
}

func main() {
  var total int = 0
  var result int = 0

  for i := 1; i < 500; i++ {
    result = fib(i)
    fmt.Printf("i: %d  Result: %d\n", i,  result)
    if result % 2 == 0 {
      total += result
    }
    if result >= 4000000 {
      fmt.Printf("Answer: %d\n", total)
      break
    }
  }

}
