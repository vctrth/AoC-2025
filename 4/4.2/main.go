package main

import (
	"fmt"
	"os"
	"strings"
)

var totalRolls int
var iterations int

func main() {
	var rolls [][]string

	content, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Fields(string(content))

	for _, row := range rows {
		temp := strings.Split(row, "")
		rolls = append(rolls, temp)
	}
	fmt.Printf("%s\n", rolls)

	//[x-1 y-1][x y-1][x+1 y-1]
	//[x-1 y]   [@]  [x+1 y]
	//[x-1 y+1][x y+1][x+1 y+1]
	startCheck(rolls, 0)
}

func startCheck(rolls [][]string, cTotal int) {
	iterations++
	var nRolls [][]string
	for y, row := range rolls {
		for x := range row {
			var coords = [2]int{y, x}
			nRolls = checkNeighbours(coords, rolls)
		}
	}
	difference := totalRolls - cTotal
	fmt.Printf("Ran iteration %d, total rolls picked up this round : %d\n", iterations, difference)
	if cTotal != totalRolls {
		startCheck(nRolls, totalRolls)
	} else {
		fmt.Printf("Total rolls you can pick up : %d\n", totalRolls)
	}

}

func checkNeighbours(aCoords [2]int, rolls [][]string) [][]string {
	var cC = [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	var count int
	found := rolls[aCoords[0]][aCoords[1]]
	for _, val := range cC {
		y_coord := aCoords[0] + val[0]
		if y_coord == -1 || y_coord >= len(rolls) {
			continue
		}
		x_coord := aCoords[1] + val[1]
		if x_coord == -1 || x_coord >= len(rolls[y_coord]) {
			continue
		}
		if rolls[y_coord][x_coord] == "@" {
			count++
		}
	}
	if count < 4 && found == "@" {
		rolls[aCoords[0]][aCoords[1]] = "."
		totalRolls++
		return rolls
	} else {
		return rolls
	}
}
