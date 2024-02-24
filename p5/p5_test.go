package p5

import "testing"

func TestPractice(t *testing.T) {
	output := Solve(1, 10)
	want := 2520
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	}
}

func TestSolve(t *testing.T) {
	output := Solve(1, 20)
	want := 232792560
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	}
}
