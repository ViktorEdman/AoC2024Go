package utils

import (
	"io"
	"os"
	"strconv"
	"strings"
)

func GetInput(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func GetSlices(input string) ([]int, []int) {
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
