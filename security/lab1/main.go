package main

import (
	"fmt"
)

type generator = func() int

func lfsr(n uint, a int, b int) generator {
	mask := 1 << (n - 1)
	return func() int {
		highestBit := 0
		if a&mask > 0 {
			highestBit = 1
		}
		a = ((a%mask)^(highestBit*b))<<1 + highestBit
		return highestBit
	}
}

func main() {
	fmt.Println("Hello, world")
}
