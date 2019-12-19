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

	part1, err := Part1(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
}

func Part1(program []int) (int, error) {
	comp := intcode.New(program)

	var count int
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			out, err := comp.Run([]int{x, y})
			if err != nil {
				return 0, err
			}
			if len(out) != 1 {
				return 0, fmt.Errorf("expected 1 output, got %d", len(out))
			}
			count += out[0]
		}
	}
	return count, nil
}
