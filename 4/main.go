package main

import (
	"fmt"
	"strings"

	"github.com/ViktorEdman/AoC2024Go/utils"
)

func main() {
	grid := utils.MakeGrid("xmxmxmxmx\nxmxmxmmxx\nxxmxmxmx")
	fmt.Println(grid)
	fmt.Println(grid[0][1])

}

func findXmas(input string) int {
	grid := utils.MakeGrid(input)
	xmasesFound := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			result := horizontalSearch(grid, i, j)
			xmasesFound += result
		}
	}
	return xmasesFound
}

func horizontalSearch(grid [][]string, y, x int) int {
	searchString := ""
	xmasesFound := 0
	for i := x - 4; i < x+4; i++ {
		if i < 0 || i > len(grid)-1 {
			continue
		}
		searchString += grid[y][i]
		grid[y][i] = "."
	}
	if strings.Contains(searchString, "XMAS") {
		fmt.Println(searchString)
		xmasesFound++
	}
	if strings.Contains(searchString, "SAMX") {
		xmasesFound++
	}
	return xmasesFound
}
