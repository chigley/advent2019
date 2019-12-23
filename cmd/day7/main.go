package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"sync"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
)

const amplifiers = 5

var (
	part1PhaseSettings = []int{0, 1, 2, 3, 4}
	part2PhaseSettings = []int{5, 6, 7, 8, 9}
)

func main() {
	program, err := advent2019.ReadIntsLine(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(program)
	if err != nil {
		log.Fatal(err)
	}

	part2, err := Part2(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Part1(program []int) (int, error) {
	comp := intcode.New(program)
	maxOutput := math.MinInt64

	for settings := range advent2019.Permutations(part1PhaseSettings) {
		output, err := outputSignal(comp, settings)
		if err != nil {
			return 0, err
		}
		maxOutput = advent2019.Max(maxOutput, output)
	}

	return maxOutput, nil
}

func Part2(program []int) (int, error) {
	maxOutput := math.MinInt64

	for settings := range advent2019.Permutations(part2PhaseSettings) {
		output, err := feedbackLoopOutputSignal(program, settings)
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

func feedbackLoopOutputSignal(program, settings []int) (int, error) {
	amps := make([]intcode.Computer, amplifiers)
	initialInput := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(amplifiers)

	input := initialInput
	for i := 0; i < amplifiers; i++ {
		amps[i] = *intcode.New(program)

		if i == amplifiers-1 {
			amps[i].RunInteractive(input, wg.Done, intcode.Outputs(initialInput))
			input <- settings[i]
		} else {
			tmp := input
			input = amps[i].RunInteractive(input, wg.Done)
			tmp <- settings[i]
		}
	}

	// Pass 0 to the first amplifier.
	initialInput <- 0

	wg.Wait()
	for _, amp := range amps {
		if err := amp.Err(); err != nil {
			return 0, err
		}
	}

	// Our result comes from our first amplifier's input channel, which is
	// equivalent to our last amplifier's output channel.
	out := <-initialInput

	return out, nil
}
