package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
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
	comp := computer{
		memory: append([]int(nil), program...),
	}

	if err := comp.write(1, noun); err != nil {
		return 0, err
	}
	if err := comp.write(2, verb); err != nil {
		return 0, err
	}

	if err := comp.run(); err != nil {
		return 0, err
	}

	return comp.read(0)
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
