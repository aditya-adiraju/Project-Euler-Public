package p6

import "testing"

func TestPractice(t *testing.T) {
	output := Solve(10)
	want := 2640
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	}
}

func TestSolve(t *testing.T) {
	output := Solve(100)
	want := 25164150
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	}
}
