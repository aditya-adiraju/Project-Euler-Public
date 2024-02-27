package p9 

import "testing"

func TestPractice(t *testing.T) {
	output := Solve(3 + 4 + 5)
	want := 3*4*5
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	}
}

func TestSolve(t *testing.T) {
	output := Solve(1000)
	want := 31875000
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	}
}
