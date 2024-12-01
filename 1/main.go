package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/ViktorEdman/AoC2024Go/utils"
)

func main() {
	input, err := utils.GetInput("input.txt")
	if err != nil {
		panic(err)
	}
	slice1, slice2 := utils.GetSlices(input)
	distance, err := countDistances(slice1, slice2)
	similarityScore := scoreSimilarities(slice1, slice2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Total distance:", distance)
	fmt.Println("Total similarity:", similarityScore)
}

func scoreSimilarities(slice1, slice2 []int) int {
	totalSimilarity := 0
	for _, num := range slice1 {
		totalOccurences := 0
		for _, num2 := range slice2 {
			if num2 == num {
				totalOccurences++
			}
		}
		totalSimilarity += (num * totalOccurences)
	}
	return totalSimilarity
}

func countDistances(slice1, slice2 []int) (int, error) {

	slices.Sort(slice1)
	slices.Sort(slice2)

	var totalDistance float64
	for i := range slice1 {
		totalDistance += math.Abs(float64(slice1[i]) - float64(slice2[i]))
	}
	return int(totalDistance), nil
}
