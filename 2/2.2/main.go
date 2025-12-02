package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main(){
	answer := 0

	//Get input
	content, err := os.ReadFile("../input.txt")
	if (err != nil){ panic(err) }
	ids := strings.Split(string(content), ",")
	var invalidIds = []string{}


	for _, value := range ids {
		value = strings.TrimSpace(value)

		// fmt.Printf("Current index: %d\n", index)
		ranges := strings.Split(value, "-")

		range_min, _ := strconv.ParseInt(ranges[0], 10, 64)
		range_max, _ := strconv.ParseInt(ranges[1], 10, 64)
		range_max ++

		fmt.Printf("Checking between %d and %d\n", range_min, range_max)
		for i := range_min; i < range_max; i++ {
			number := strconv.Itoa(int(i))

			//Chunks
			for c := 1; c <= (len(number)/2); c ++ {
				chunk := number[:c]
				repeats := len(number) / c
				temp := strings.Repeat(chunk, repeats)
				fmt.Printf("checking if %s = %s\n", temp, number)
				if temp == number && !contains(invalidIds, number){
					invalidIds = append(invalidIds, number)
				}
			}
		}
		fmt.Printf("Invalid Ids in range : %s\n", invalidIds)
	}

	//Calculate answer
	for i := 0; i < len(invalidIds); i++ {
		n, err := strconv.Atoi(invalidIds[int(i)])
		if err != nil { panic(err) }
		answer += n
	}

	fmt.Printf("The answer is : %d\n", answer)
}

func contains(array []string, element string) bool{
	for i := 0; i < len(array); i++ {
		if(array[i] == element) { return true }
	}
	return false
}