package p3 

import "testing"

func TestPractice(t *testing.T) {
	output := Solve(13195)
	want := 29
	if output != want {
		t.Fatalf("Expected: %#v \n Got: %#v", want, output)
	}
}

func TestSolve(t *testing.T) {
	output := Solve(600851475143)
	want := 6857
	if output != want {
		t.Fatalf("Expected: %#v \n Got: %#v", want, output)
	}
}
