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
