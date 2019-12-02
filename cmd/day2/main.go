package main

import (
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
	fmt.Println(part1)
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
