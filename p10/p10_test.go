package p10

import "testing"

func TestPractice(t *testing.T) {
	output := Solve(10)
	want := 17
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	} 

}


func TestSolve(t *testing.T) {
	output := Solve(2000000)
	want := 142913828922
	if output != want {
		t.Fatalf("Expected: %#v\nGot: %#v", want, output)
	} 

}
