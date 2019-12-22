package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"

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
	// Jumps take 4 steps. Jump when there isn't a hole to worry about in the
	// next 3 spaces, and it's safe to land.
	//
	// !(A && B && C) && D
	script := `OR A T
AND B T
AND C T
NOT T J
AND D J
WALK
`
	return attempt(comp, script)
}

func Part2(comp *intcode.Computer) (int, error) {
	// Jump if the part 1 condition is true and we will also be able to safely
	// walk or jump from where we land.
	//
	// !(A && B && C) && D && (E || H)
	script := `OR A T
AND B T
AND C T
NOT T T
AND D T
OR E J
OR H J
AND T J
RUN
`
	return attempt(comp, script)
}

func attempt(comp *intcode.Computer, script string) (int, error) {
	outputs, err := comp.Run(intcode.ToASCII(script))
	if err != nil {
		return 0, err
	}
	if len(outputs) == 0 {
		return 0, errors.New("got no output")
	}

	last := outputs[len(outputs)-1]
	if last > unicode.MaxASCII {
		return last, nil
	}
	return 0, errors.New(intcode.FromASCII(outputs))
}
