package main

import "testing"

var input = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestFind(t *testing.T) {
	want := 18
	got := findXmas(input)
	if got != want {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
