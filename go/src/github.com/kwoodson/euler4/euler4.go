package main
// a palindromic number

import (
	"fmt"
        "strconv"
)

func ispalindromic(n int) bool {
  int_str := strconv.Itoa(n)
  var str_len int = len(int_str) - 1

  for i := 0; i < str_len; i++ {
    if (int_str[i] != int_str[str_len - i]) {
      return false
    }
  }
  return true
}

func main() {
  var z int = 0
  var ans int = 0
  for i := 999; i > 0; i-- {
    for j := 999; j >= i; j-- {
      z = i * j
      if ispalindromic(z) {
        //fmt.Printf("palindrome: %d\n", z)
        if (ans < z) {
          ans = z
        }
      }
    }
  }
  fmt.Printf("highest palindrome: %d\n", ans)
}
