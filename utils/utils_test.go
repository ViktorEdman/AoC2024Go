package utils

import (
	"reflect"
	"testing"
)

var input = `3 3 3 4
5 5 5 3
20 88 333 44`

func TestGetIntSlices(t *testing.T) {
	want := [][]int{
		{3, 3, 3, 4},
		{5, 5, 5, 3},
		{20, 88, 333, 44},
	}
	got, err := GetIntSlices(input)
	if err != nil {
		t.Fatalf("Got error %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)

	}
}
