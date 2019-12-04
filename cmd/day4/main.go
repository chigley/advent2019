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
	input := ranges[0]

	fmt.Println(part1(input))
}

func part1(input advent2019.Range) (validPasswords int) {
	for i := input.From; i <= input.To; i++ {
		if meetsCriteria(i) {
			validPasswords++
		}
	}
	return
}

func meetsCriteria(x int) bool {
	code := fmt.Sprintf("%06d", x)

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
