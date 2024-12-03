package main

import (
	"fmt"
	"regexp"

	"github.com/ViktorEdman/AoC2024Go/utils"
)

func main() {
	input, err := utils.GetInput("input.txt")
	if err != nil {
		panic(err)
	}
	muls := findMuls(input)
	sum, err := execMuls(muls)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result of muls: ", sum)

}

func findMuls(input string) []string {
	pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	muls := pattern.FindAllString(input, -1)
	return muls
}

func execMuls(muls []string) (int, error) {
	sum := 0
	for _, mul := range muls {
		ints, err := utils.GetInts(mul)
		if err != nil {
			return 0, err
		}
		product := 1
		for _, num := range ints {
			product *= num
		}
		sum += product
		fmt.Println(mul, ints, product, sum)
	}
	return sum, nil
}
