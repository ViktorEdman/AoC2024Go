package main

import "testing"

var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestVerifySafe(t *testing.T) {
	want := 2
	got := verifySafe(getIntSlices(input))
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
