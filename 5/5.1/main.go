package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ranges []string
	var ids []int
	var fresh []int
	content, _ := os.ReadFile("../input.txt")
	// fmt.Printf("%s\n", string(content))
	t := strings.FieldsSeq(string(content))
	for subs := range t {
		if strings.Contains(subs, "-") {
			ranges = append(ranges, subs)
		} else {
			id, _ := strconv.Atoi(subs)
			ids = append(ids, id)
		}
	}
	fmt.Println("Input parsed!")
	for _, id := range ids {
		if checkFresh(id, ranges) {
			fresh = append(fresh, id)
			fmt.Println("Fresh ingredient found")
		}
	}
	fmt.Printf("%d fresh ingredients found\n", len(fresh))
}

func checkFresh(id int, arr []string) bool {
	for _, a := range arr {
		ran := strings.Split(a, "-")
		min, _ := strconv.Atoi(ran[0])
		max, _ := strconv.Atoi(ran[1])
		fmt.Printf("Checking from %d to %d\n", min, max)
		if min <= id && id <= max {
			return true
		}
	}
	return false
}
