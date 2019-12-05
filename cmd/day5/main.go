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

	outputs, err := intcode.New(program).Run([]int{1})
	if err != nil {
		log.Fatal(err)
	}
	for i, out := range outputs {
		if i != len(outputs)-1 && out != 0 {
			log.Fatal("got a non-zero output before the end")
		}
	}
	fmt.Println(outputs[len(outputs)-1])
}
