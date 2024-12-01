package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/ViktorEdman/AoC2024Go/utils"
)

func main() {
	input, err := utils.GetInput("input.txt")
	if err != nil {
		panic(err)
	}
	slice1, slice2 := getSlices(input)
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

func getSlices(input string) ([]int, []int) {
	slice1 := []int{}
	slice2 := []int{}
	list := strings.Split(input, "\n")
	for _, row := range list {
		if len(row) < 2 {
			break
		}
		pair := strings.Split(row, "   ")
		num1, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		slice1 = append(slice1, num1)
		slice2 = append(slice2, num2)
	}
	return slice1, slice2
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
