package main

import (
	"fmt"
)

func checkSeq(s []int) int {
	n := len(s)
	L := 0
	m := -1
	b := make([]int, n)
	c := make([]int, n)
	t := make([]int, n)
	b[0] = 1
	c[0] = 1
	for N := 0; N < n; N++ {
		d := s[N]
		for i := 1; i <= L; i++ {
			d = d ^ (c[i] * s[N-i])
		}
		if d != 0 {
			copy(t, c)
			for i := 0; i < n-N+m; i++ {
				c[N-m+i] = c[N-m+i] ^ b[i]
			}
			if 2*L <= N {
				L = N + 1 - L
				m = N
				copy(b, t)
			}
		}
	}
	return L
}

func main() {
	gen := labGen()
	seq := make([]int, 40)
	fmt.Printf("Sequence: ")
	for i := range seq {
		seq[i] = gen()
		fmt.Printf("%d, ", seq[i])
	}
	fmt.Println("end")
	fmt.Printf("Linear complexity: %d\n", checkSeq(seq))
}
