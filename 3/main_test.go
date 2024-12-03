package main

import (
	"reflect"
	"testing"
)

var input = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
var testMuls = []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}

func TestExecMuls(t *testing.T) {
	want := 161
	got, err := execMuls(testMuls)
	if err != nil {
		t.Fatalf("Got error: %v", err)
	}
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestFindMuls(t *testing.T) {
	want := testMuls
	got := findMuls(input)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
