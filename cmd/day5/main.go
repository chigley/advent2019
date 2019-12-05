package main

import (
	"errors"
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

	part1, err := computePart1(comp)
	if err != nil {
		log.Fatal(err)
	}

	part2, err := computePart2(comp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func computePart1(comp *intcode.Computer) (int, error) {
	outputs, err := comp.Run([]int{1})
	if err != nil {
		log.Fatal(err)
	}
	if len(outputs) == 0 {
		return 0, errors.New("got no outputs")
	}
	for i, out := range outputs {
		if i != len(outputs)-1 && out != 0 {
			return 0, errors.New("got a non-zero output before the end")
		}
	}
	return outputs[len(outputs)-1], nil
}

func computePart2(comp *intcode.Computer) (int, error) {
	outputs, err := comp.Run([]int{5})
	if err != nil {
		return 0, err
	}
	if len(outputs) != 1 {
		return 0, fmt.Errorf("expected exactly one output, got %d", len(outputs))
	}
	return outputs[0], nil
}
