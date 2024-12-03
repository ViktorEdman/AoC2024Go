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
	muls = findMulsAndDos(input)
	sum, err = execMuls(muls)
	if err != nil {
		panic(err)
	}
	fmt.Println("Result of enabled muls: ", sum)

}

var mulPattern = `mul\([0-9]+,[0-9]+\)`
var doPattern = `do(n't)*\(\)`

func findMuls(input string) []string {
	pattern := regexp.MustCompile(mulPattern)
	muls := pattern.FindAllString(input, -1)
	return muls
}

func findMulsAndDos(input string) []string {
	pattern := regexp.MustCompile(fmt.Sprintf(`(%s)|(%s)`, mulPattern, doPattern))
	muls := pattern.FindAllString(input, -1)
	return muls
}

func execMuls(muls []string) (int, error) {
	sum := 0
	enabled := true
	for _, mul := range muls {
		if mul == "do()" {
			enabled = true
			continue
		}
		if mul == "don't()" {
			enabled = false
			continue
		}
		if enabled == false {
			continue
		}
		ints, err := utils.GetInts(mul)
		if err != nil {
			return 0, err
		}
		product := 1
		for _, num := range ints {
			product *= num
		}
		sum += product
	}
	return sum, nil
}
