package main

import (
	"fmt"
	"testing"

	"github.com/ViktorEdman/AoC2024Go/utils"
)

var input = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestVerifySafe(t *testing.T) {
	want := 2
	got := 0
	reports, _ := utils.GetIntSlices(input)
	for _, report := range reports {
		if verifySafe(report) {
			fmt.Println(report, "undampened safe")
			got++
			continue
		}
		fmt.Println(report, "undampened unsafe")
	}
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}

func TestDampenedVerifySafe(t *testing.T) {
	want := 4
	got := 0
	reports, _ := utils.GetIntSlices(input)
	for _, report := range reports {
		if dampenedVerifySafe(report) {
			fmt.Println(report, "safe")
			got++
			continue
		}
		fmt.Println(report, "unsafe")
	}
	if want != got {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
