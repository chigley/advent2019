package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
)

func main() {
	program, err := advent2019.ReadIntsLine(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	comp := intcode.New(program)

	part1, err := Part1(comp)
	if err != nil {
		log.Fatal(err)
	}

	part2, err := Part2(comp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Part1(comp *intcode.Computer) (int, error) {
	outputs, err := comp.Run([]int{1})
	if err != nil {
		return 0, err
	}
	if len(outputs) != 1 {
		return 0, fmt.Errorf("expected 1 output, got %d", len(outputs))
	}
	return outputs[0], nil
}

func Part2(comp *intcode.Computer) (int, error) {
	outputs, err := comp.Run([]int{2})
	if err != nil {
		return 0, err
	}
	if len(outputs) != 1 {
		return 0, fmt.Errorf("expected 1 output, got %d", len(outputs))
	}
	return outputs[0], nil
}
