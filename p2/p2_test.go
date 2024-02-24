package p2

import (
	"testing"
)

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}	
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true 
}

func TestPractice(t *testing.T) {
	output := Practice()
	want := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	if !equal(output, want) {
		t.Fatalf("Expected: %#v,\n Got: %#v", want, output)
	}
}

func TestSolve(t *testing.T) {
	output := Solve(4000000)
	want := 4613732
	if output != want {
		t.Fatalf("Expected: %#v,\n Got: %#v", want, output)
	}
}
