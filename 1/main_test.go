package main

import (
	"testing"

	"github.com/ViktorEdman/AoC2024Go/utils"
)

var input = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestCountDistances(t *testing.T) {
	want := 11
	result, err := countDistances(utils.GetSlices(input))
	if err != nil {
		t.Error(err)
	}
	if want != result {
		t.Errorf("Wanted %v, got %v", want, result)
	}
}

func TestScoreSimilarities(t *testing.T) {
	want := 31
	result := scoreSimilarities(utils.GetSlices(input))
	if want != result {
		t.Errorf("Wanted %v, got %v", want, result)
	}
}
