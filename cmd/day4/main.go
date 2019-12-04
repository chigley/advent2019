package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
)

func main() {
	ranges, err := advent2019.ReadRanges(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if len(ranges) != 1 {
		log.Fatal("expected one input range")
	}

	part1, part2 := solve(ranges[0])
	fmt.Println(part1)
	fmt.Println(part2)
}

func solve(input advent2019.Range) (part1, part2 int) {
	for i := input.From; i <= input.To; i++ {
		code := fmt.Sprintf("%06d", i)
		if meetsPart1Criteria(code) {
			part1++

			if meetsPart2AdditionalCriteria(code) {
				part2++
			}
		}
	}
	return
}

func meetsPart1Criteria(code string) bool {
	if !advent2019.StringIsSorted(code) {
		return false
	}

	runes := []rune(code)
	for i, r := range runes {
		if i == 0 {
			continue
		}
		if r == runes[i-1] {
			return true
		}
	}

	return false
}

func meetsPart2Criteria(code string) bool {
	return meetsPart1Criteria(code) && meetsPart2AdditionalCriteria(code)
}

func meetsPart2AdditionalCriteria(code string) bool {
	var count int

	runes := []rune(code)
	for i, r := range runes {
		if i > 0 && r == runes[i-1] {
			count++

			if i == len(runes)-1 && count == 2 {
				return true
			}
		} else {
			if count == 2 {
				return true
			}
			count = 1
		}
	}

	return false
}
