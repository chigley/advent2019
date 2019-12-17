package main

import (
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

	part1, err := Part1(input[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
}

func Part1(input string) (string, error) {
	digits, err := parseDigits(input)
	if err != nil {
		return "", err
	}
	return first8(phaseN(digits, 100))
}

func phaseN(input []int, n int) []int {
	for i := 0; i < n; i++ {
		input = phase(input)
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

func parseDigits(input string) ([]int, error) {
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
