package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 1.0; a < 1000; a++ {
		for b := a + 1.0; b < 1000; b++ {

			c := a*a + b*b
			sqrtC := math.Sqrt(c)
			if sqrtC > a && sqrtC > b && a+b+sqrtC == 1000 {
				fmt.Printf("Found it: a=%v b=%v c=%v\n", a, b, sqrtC)

			}
		}
	}
}
