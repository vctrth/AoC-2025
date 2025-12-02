package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

var dial int
var password int
var values []string
var turns int
var direction int

func main() {
	dial = 50

	content, err := os.ReadFile("../input.txt")
	if err != nil { panic(err) }
	values = strings.Fields(string(content))

	for i := 0; i< len(values); i++ {
		turn(calcTurn(values[i]))
	}

	fmt.Println("Dial currently on ", dial)
	fmt.Println("password = ", password)
	fmt.Println("total nubers of turns = ", turns)
}

func calcTurn(value string) int{
	if (strings.Contains(value, "L")){
		//Turn left
		after, found := strings.CutPrefix(value, "L")
		if (!found) { return 0 }

		val, err := strconv.ParseInt(after, 10, 16)
		if err != nil { panic(err) }

		direction = -1
		return int(val)
	} else if(strings.Contains(value, "R")) {
		//Turn right
		after, found := strings.CutPrefix(value, "R")
		if (!found) { return 0 }

		val, err := strconv.ParseInt(after, 10, 16)
		if err != nil { panic(err) }

		direction = 1
		return int(val)
	}
	return 0
}

func checkTurn(){
	if dial == 100 { 
		dial = 0 
	} else if dial == -1 { 
		dial = 99 
	}
}

func turn(value int){
	turns ++
	fmt.Println("will turn dial with ", value)
	for i := 0; i<value; i++ {
		dial += 1 * direction
		checkTurn()
		if dial == 0 { password ++ }
	}
	fmt.Println("Dial posistion after turning = ", dial)
}