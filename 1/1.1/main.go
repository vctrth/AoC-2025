package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dial int
var password int
var values []string
var turns int

func main() {
	dial = 50

	content, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	values = strings.Fields(string(content))

	for i := 0; i < len(values); i++ {
		turn(calcTurn(values[i]))
	}

	fmt.Println("password = ", password)
	fmt.Println("total nubers of turns = ", turns)
}

func calcTurn(value string) int {
	if strings.Contains(value, "L") {
		//Turn left
		after, found := strings.CutPrefix(value, "L")
		if !found {
			return 0
		}

		val, err := strconv.ParseInt(after, 10, 16)
		if err != nil {
			panic(err)
		}
		return -int(val)
	} else if strings.Contains(value, "R") {
		//Turn right
		after, found := strings.CutPrefix(value, "R")
		if !found {
			return 0
		}

		val, err := strconv.ParseInt(after, 10, 16)
		if err != nil {
			panic(err)
		}
		return int(val)
	}
	return 0
}

func turn(value int) {
	turns++
	fmt.Println("will turn dial with ", value)
	dial += value

	for dial > 99 {
		dial -= 100
	}
	for dial < 0 {
		dial += 100
	}

	fmt.Println("dial currently on ", dial)
	if dial == 0 {
		password++
	}
}
