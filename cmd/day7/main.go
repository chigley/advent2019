package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
)

const amplifiers = 5

var phaseSettings = []int{0, 1, 2, 3, 4}

func main() {
	program, err := advent2019.ReadIntsLine(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
}

func Part1(program []int) (int, error) {
	comp := intcode.New(program)
	maxOutput := math.MinInt64

	for settings := range advent2019.Permutations(phaseSettings) {
		output, err := outputSignal(comp, settings)
		if err != nil {
			return 0, err
		}
		maxOutput = advent2019.Max(maxOutput, output)
	}

	return maxOutput, nil
}

func outputSignal(comp *intcode.Computer, settings []int) (int, error) {
	var input int
	for i := 0; i < amplifiers; i++ {
		out, err := comp.Run([]int{settings[i], input})
		if err != nil {
			return 0, err
		}
		if len(out) != 1 {
			return 0, errors.New("expected exactly one output")
		}
		input = out[0]
	}

	return input, nil
}
