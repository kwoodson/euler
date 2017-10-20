package main
// The sum of the squares of the first ten natural numbers is:
// 1^2 + 2^2 + .. 10^2 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + 3 + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square
// of the sum is 3025-385 = 2640
// Find the difference between the sum of the squares of the first one hundred natrual numbers and the square of the sum.


import (
	"fmt"
)

func iPow(a, b int64) int64 {
  var result int64 = 1;

  for 0 != b {
    if 0 != (b & 1) {
      result *= a;

    }
    b >>= 1;
    a *= a;
  }

  return result;
}

func main() {
  var sum_square int64 = 0
  var square_sum int64 = 0

  for i := int64(1); i <= 100; i++ {
    sum_square += iPow(i, 2)//i**i
    square_sum += i
  }
  fmt.Printf("Prime: %d\n", (iPow(square_sum, 2) - sum_square))
}
