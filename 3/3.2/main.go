package 
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var combs []string
	content, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	banks := strings.Fields(string(content))
	fmt.Printf("Loaded in %s\n", banks)
	for _, bank := range banks {
		a := analyzeBank(bank)
		fmt.Printf("The highest combination is : %s\n", a)
		combs = append(combs, a)
	}
	fmt.Printf("The answer is : %d\n", calculateAnswer(combs))
}

func analyzeBank(bank string) string {
	var highest_left int
	var highest_left_i int
	var highest_right int
	for i := 0; i < (len(bank) - 1); i++ {
		digit, err := strconv.Atoi(string(bank[i]))
		if err != nil {
			panic(err)
		}
		if digit > highest_left {
			highest_left = digit
			highest_left_i = i
		}
	}
	for i := (len(bank) - 1); i > highest_left_i; i-- {
		digit, err := strconv.Atoi(string(bank[i]))
		if err != nil {
			panic(err)
		}
		if digit > highest_right {
			highest_right = digit
		}
	}
	return fmt.Sprintf("%d%d", highest_left, highest_right)
}

func calculateAnswer(combs []string) int {
	var total int
	for _, comb := range combs {
		combI, _ := strconv.Atoi(comb)
		total += combI
	}
	return total
}
