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

	droid := NewDroid(program)

	part1, err := droid.Part1()
	if err != nil {
		log.Fatal(err)
	}

	part2, err := droid.Part2()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func (d *Droid) Part1() (int, error) {
	return d.bfs(true)
}

func (d *Droid) Part2() (int, error) {
	return d.bfs(false)
}
