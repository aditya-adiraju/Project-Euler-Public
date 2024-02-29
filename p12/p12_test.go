package p12

import "testing"

func TestPractice(t *testing.T) {
	output := Solve(5)
	want := 28
	if output != want {
		t.Fatalf("Output: %#v", output)
	}
}

func TestSolve(t *testing.T) {
	output := Solve(500)
	t.Logf("Result: %#v", output)
	
}
