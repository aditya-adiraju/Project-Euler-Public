package p14

import "testing"


func TestPractice(t *testing.T) {
	output := Solve(10)
	want := 9
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	}
}

func TestSolve(t *testing.T) {
	output := Solve(1000000)
	t.Logf("Result: %d", output)
}
