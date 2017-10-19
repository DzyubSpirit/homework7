package main

import "testing"

func TestLFSR(t *testing.T) {
	lfsr1 := lfsr(4, 12, 5)
	expectedOutputs := []int{1, 0, 0, 1, 0, 0}
	for i, expectedOutput := range expectedOutputs {
		output := lfsr1()
		if output != expectedOutput {
			t.Errorf("Output #%d failed, expected: %d, got: %d", i+1, expectedOutput, output)
		}
	}
}

func TestBinaryCombine(t *testing.T) {
	binaryGen := func(count int) generator {
		val := 0
		i := 0
		return func() int {
			if i == count {
				i = 0
				val = 1 - val
			}
			i++
			return val
		}
	}
	gen := binaryCombine(binaryGen(4), binaryGen(2), binaryGen(1))
	expectedOutputs := []int{0, 1, 2, 3, 4, 5, 6, 7}
	for i, expectedOutput := range expectedOutputs {
		output := gen()
		if output != expectedOutput {
			t.Errorf("Output #%d failed, expected: %d, got: %d", i+1, expectedOutput, output)
		}
	}

}

func TestCheckSeq(t *testing.T) {
	seq := []int{1, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1}
	output := checkSeq(seq)
	if output != 4 {
		t.Errorf("CheckSeq failed, expected: %d, got: %d", 4, output)

	}
}
