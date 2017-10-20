package main

import "math/rand"

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

func binaryCombine(gens ...generator) generator {
	return func() int {
		position := 0
		for _, gen := range gens {
			position = position<<1 + gen()
		}
		return position
	}
}

func tableGen(gen generator, table []int) generator {
	return func() int {
		return table[gen()]
	}
}

func labGen() generator {
	const n = 10
	lfrsArr := make([]generator, n)
	for i, p := range [n]struct {
		n uint
		a int
		b int
	}{
		{8, 142, 142},
		{9, 264, 264},
		{10, 516, 516},
		{11, 1026, 1026},
		{12, 2089, 2089},
		{13, 4109, 4109},
		{14, 8737, 8737},
		{15, 16385, 16385},
		{16, 34821, 34821},
		{17, 65540, 65540},
	} {
		lfrsArr[i] = lfsr(p.n, p.a, p.b)
	}

	const tableSize = 1024
	var table [tableSize]int
	for i := 0; i < tableSize/2; i++ {
		position := rand.Int() % tableSize
		for table[position] == 1 {
			position = rand.Int() % tableSize
		}
		table[position] = 1
	}

	return tableGen(binaryCombine(lfrsArr...), table[:])
}
