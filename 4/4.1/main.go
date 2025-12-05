package main

import (
	"fmt"
	"os"
	"strings"
)

var totalRolls int

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
	for y, row := range rolls {
		for x := range row {
			var coords = [2]int{y, x}
			f, c, b := checkNeighbours(coords, rolls)
			fmt.Printf("object : %s, neighbours? : %d, pickup? : %t\n", f, c, b)
		}
	}
	fmt.Printf("Total rolls you can pick up : %d\n", totalRolls)
}

func checkNeighbours(aCoords [2]int, rolls [][]string) (string, int, bool) {
	var cC = [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	var count int
	found := rolls[aCoords[0]][aCoords[1]]
	if found != "@" {
		return found, count, false
	}
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
	if count < 4 {
		totalRolls++
		return found, count, true
	} else {
		return found, count, false
	}
}
