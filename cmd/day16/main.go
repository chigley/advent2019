package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chigley/advent2019"
)

var basePattern = []int{0, 1, 0, -1}

func main() {
	input, err := advent2019.ReadStrings(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	if len(input) != 1 {
		log.Fatal("expected one input line")
	}

	digits, err := ParseDigits(input[0])
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(digits)
	if err != nil {
		log.Fatal(err)
	}

	part2, err := Part2(input[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Part1(input []int) (string, error) {
	return first8(phaseN(input, phase, 100))
}

func Part2(input string) (string, error) {
	offset, err := strconv.Atoi(input[:7])
	if err != nil {
		return "", err
	}

	repeatedInputLen := len(input) * 10000
	if offset < (repeatedInputLen+1)/2 {
		return "", errors.New("offset too low for our optimisation to work")
	}

	input = strings.Repeat(input, 10000)

	// Optimisation one: each digit does not depend on any digits before it. We
	// can throw away everything before the offset without affecting our result
	input = input[offset:]

	digits, err := ParseDigits(input)
	if err != nil {
		return "", err
	}

	return first8(phaseN(digits, phaseOptimised, 100))
}

func phaseN(input []int, f func([]int) []int, n int) []int {
	for i := 0; i < n; i++ {
		input = f(input)
	}
	return input
}

func phase(input []int) []int {
	output := make([]int, len(input))
	for i := range input {
		var sum int
		for j, digit := range input {
			sum += digit * basePattern[((j+1)/(i+1))%len(basePattern)]
		}
		output[i] = advent2019.Abs(sum % 10)
	}
	return output
}

// Optimisation two: from the middle of the overall input (_not_ the input to
// this function) onwards, the next value of a given digit can be derived as its
// current value plus the new values of all subsequent digits.
//
// This only makes sense if the overall input string has been truncated at the
// midpoint or later.
func phaseOptimised(input []int) []int {
	output := make([]int, len(input))
	for i := len(output) - 1; i >= 0; i-- {
		if i == len(output)-1 {
			output[i] = input[i]
			continue
		}
		output[i] = (input[i] + output[i+1]) % 10
	}
	return output
}

func ParseDigits(input string) ([]int, error) {
	digits := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		digit, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return nil, err
		}
		digits[i] = digit
	}
	return digits, nil
}

func first8(digits []int) (string, error) {
	if len(digits) < 8 {
		return "", fmt.Errorf("got %d digits, expected 8+", len(digits))
	}

	var b strings.Builder
	for _, n := range digits[:8] {
		b.WriteString(strconv.Itoa(n))
	}
	return b.String(), nil
}
