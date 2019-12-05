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

	outputs, err := comp.Run([]int{1})
	if err != nil {
		log.Fatal(err)
	}
	for i, out := range outputs {
		if i != len(outputs)-1 && out != 0 {
			log.Fatal("got a non-zero output before the end")
		}
	}
	part1 := outputs[len(outputs)-1]

	outputs, err = comp.Run([]int{5})
	if err != nil {
		log.Fatal(err)
	}
	if len(outputs) != 1 {
		log.Fatal("expected exactly one output")
	}
	part2 := outputs[0]

	fmt.Println(part1)
	fmt.Println(part2)

}
