package main

import "testing"

var input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestCountDistances(t *testing.T) {
	want := 11
	result, err := countDistances(getSlices(input))
	if err != nil {
		t.Error(err)
	}
	if want != result {
		t.Errorf("Wanted %v, got %v", want, result)
	}
}

func TestScoreSimilarities(t *testing.T) {
	want := 31
	result := scoreSimilarities(getSlices(input))
	if want != result {
		t.Errorf("Wanted %v, got %v", want, result)
	}
}
