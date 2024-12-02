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

func GetInts(input string) ([]int, error) {
	currentNumString := ""
	ints := []int{}
	for _, rune := range input {
		if rune <= '9' && rune >= '0' {
			currentNumString += string(rune)
			continue
		} else if currentNumString != "" {
			int, err := strconv.Atoi(currentNumString)
			if err != nil {
				return nil, err

			}
			ints = append(ints, int)
			currentNumString = ""
		}
	}
	if currentNumString != "" {
		int, err := strconv.Atoi(currentNumString)
		if err != nil {
			return nil, err
		}
		ints = append(ints, int)
	}
	return ints, nil
}

func GetIntSlices(input string) ([][]int, error) {
	rows := strings.Split(input, "\n")
	slices := [][]int{}
	for _, row := range rows {
		slice, err := GetInts(row)
		if err != nil {
			return nil, err
		}
		slices = append(slices, slice)
	}
	return slices, nil
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
