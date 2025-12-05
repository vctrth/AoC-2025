// Old version that expanded the ranges
package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func archive() {
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
	fmt.Println("Started conversion")
	iR := convRanges(ranges)
	for _, id := range ids {
		if slices.Contains(iR, id) {
			fresh = append(fresh, id)
		}
	}
	fmt.Printf("%d fresh ingredients found : %d\n", len(fresh), fresh)
}

func convRanges(arr []string) []int {
	var r []int
	for _, a := range arr {
		// r = fmt.Sprintf("%s%s", r, a)
		ran := strings.Split(a, "-")
		min, _ := strconv.Atoi(ran[0])
		max, _ := strconv.Atoi(ran[1])
		for i := min; i <= max; i++ {
			r = append(r, i)
		}
	}
	fmt.Println("Ranges converted!")
	return r
}
