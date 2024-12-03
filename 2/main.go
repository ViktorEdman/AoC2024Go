package main

import (
	"fmt"
	"math"

	"github.com/ViktorEdman/AoC2024Go/utils"
)

func main() {
	input, err := utils.GetInput("input.txt")
	if err != nil {
		panic(err)
	}
	reports, err := utils.GetIntSlices(input)
	if err != nil {
		panic(err)
	}
	safe := 0
	dampenedSafe := 0
	for _, report := range reports {
		if verifySafe(report) {
			safe++
		}
		if dampenedVerifySafe(report) {
			dampenedSafe++
		}
	}
	fmt.Println("Safe reports:", safe)
	fmt.Println("Dampened safe reports:", dampenedSafe)

}

func dampenedVerifySafe(report []int) (safe bool) {
	if verifySafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		dampenedReport := make([]int, 0)
		for j := 0; j < len(report); j++ {
			if j == i {
				continue
			}
			dampenedReport = append(dampenedReport, report[j])
		}
		if verifySafe(dampenedReport) {
			return true
		}
	}
	return false
}

func verifySafe(report []int) (safe bool) {
	increasing := 0
	decreasing := 0
	for i := range report {
		if i == 0 {
			continue
		}
		lastNum := report[i-1]
		currentNum := report[i]
		diff := currentNum - lastNum
		if math.Abs(float64(diff)) > 3 || math.Abs(float64(diff)) < 1 {
			return false
		}
		if currentNum == lastNum {
			return false
		}
		if currentNum > lastNum {
			increasing++
		}
		if currentNum < lastNum {
			decreasing++
		}

	}
	if increasing == len(report)-1 || decreasing == len(report)-1 {
		return true
	}
	return false

}
