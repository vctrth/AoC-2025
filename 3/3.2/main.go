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
		a := analyseBank(bank, 12)
		fmt.Printf("The highest combination is : %s\n", a)
		combs = append(combs, a)
	}
	fmt.Printf("The answer is : %d\n", calculateAnswer(combs))
}

func analyseBank(bank string, nB int) string {
	var values []string
	var highest_i int
	for n := 1; n <= nB; n++ {
		if n == 1 {
			fmt.Printf("Checking for first digit\n")
			var highest int
			for i := 0; i < (len(bank) - (nB - n)); i++ {
				digit, err := strconv.Atoi(string(bank[i]))
				if err != nil {
					panic(err)
				}
				if digit > highest {
					highest = digit
					highest_i = i
					fmt.Printf("First highest index : %d\n", highest_i)
				}
			}
			values = append(values, strconv.Itoa(highest))
		} else {
			// fmt.Printf("Checking for digit %d, max index : %d\n", n, (len(bank) - (nB - n)))
			// fmt.Printf("Starting at %d\n", (highest_i + 1))
			var highest int
			for i := (highest_i + 1); i < (len(bank) - (nB - n)); i++ {
				digit, err := strconv.Atoi(string(bank[i]))
				if err != nil {
					panic(err)
				}
				// fmt.Printf("Checking digit %d\n", digit)
				if digit > highest {
					highest = digit
					highest_i = i
					fmt.Printf("New highest index : %d\n", highest_i)
				}
			}
			values = append(values, strconv.Itoa(highest))
		}
		fmt.Printf("values : %s\n", values)
	}
	return strings.Join(values, "")
}

func calculateAnswer(combs []string) int {
	var total int
	for _, comb := range combs {
		combI, _ := strconv.Atoi(comb)
		total += combI
	}
	return total
}
