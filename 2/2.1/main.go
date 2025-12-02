package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main(){
	answer := 0

	content, err := os.ReadFile("../input.txt")
	if (err != nil){ panic(err) }
	ids := strings.Split(string(content), ",")
	var invalidIds = []string{}

	for _, value := range ids {

		value = strings.TrimSpace(value)

		// fmt.Printf("Current index: %d\n", index)
		ranges := strings.Split(value, "-")

		range_min, err := strconv.ParseInt(ranges[0], 10, 64)
		if err != nil { panic(err) }

		range_max, err := strconv.ParseInt(ranges[1], 10, 64)
		if err != nil { panic(err) }
		range_max ++

		fmt.Printf("Checking between %d and %d\n", range_min, range_max)
		for i := range_min; i < range_max; i++ {
			//Max repetition length
			max := len(strconv.Itoa(int(i)))
			number := strconv.Itoa(int(i))
			half := max / 2
			left := number[:half]
			right := number[half:]
			// fmt.Printf("left: %s right: %s\n", left, right)
			if(left == right) {
				invalidIds = append(invalidIds, number)
			}
		}
		fmt.Printf("Invalid Ids in range : %s\n", invalidIds)
	}
	for i := 0; i < len(invalidIds); i++ {
		n, err := strconv.Atoi(invalidIds[int(i)])
		if err != nil { panic(err) }
		// fmt.Printf("Adding %d to %d\n", n, answer)
		answer += n
	}

	fmt.Printf("The answer is : %d\n", answer)
}