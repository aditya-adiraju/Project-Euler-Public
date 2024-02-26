package p7

import "testing"

func TestSolve(t *testing.T) {
	output := Solve(10001)
	t.Logf("Result: %#v", output)
} 
