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

	part1, err := evaluate(program, 12, 2)
	if err != nil {
		log.Fatal(err)
	}

	noun, verb, err := findInputs(program, 19690720)
	if err != nil {
		log.Fatal(err)
	}
	part2 := 100*noun + verb

	fmt.Println(part1)
	fmt.Println(part2)
}

func evaluate(program []int, noun, verb int) (int, error) {
	if len(program) < 3 {
		return 0, errors.New("program not long enough to write noun and verb")
	}
	program[1] = noun
	program[2] = verb

	comp := intcode.New(program)

	if err := comp.Run(); err != nil {
		return 0, err
	}

	return comp.Read(0)
}

func findInputs(program []int, target int) (int, int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result, err := evaluate(program, noun, verb)
			if err != nil {
				return 0, 0, err
			}
			if result == target {
				return noun, verb, nil
			}
		}
	}
	return 0, 0, errors.New("no suitable noun/verb combination found")
}
