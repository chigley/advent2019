package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
	"github.com/chigley/advent2019/vector"
)

type screen map[vector.XY]tile

type tile int

const (
	tileEmpty tile = iota
	tileWall
	tileBlock
	tilePaddle
	tileBall
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

	part2, err := Part2(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Part1(program []int) (int, error) {
	out, err := intcode.New(program).Run(nil)
	if err != nil {
		return 0, err
	}

	screen := make(screen)
	for i := 0; i < len(out); i += 3 {
		x, y, tileID := out[i], out[i+1], out[i+2]
		screen[vector.XY{X: x, Y: y}] = tile(tileID)
	}

	var blocks int
	for _, tile := range screen {
		if tile == tileBlock {
			blocks++
		}
	}
	return blocks, nil
}

func Part2(program []int) (int, error) {
	if len(program) == 0 {
		return 0, errors.New("program too short")
	}
	program[0] = 2
	comp := intcode.New(program)

	inputs := make(chan int)
	done := make(chan struct{})
	outputs := comp.RunInteractive(inputs, func() {
		close(done)
	})

	var (
		score   int
		ballX   int
		paddleX int
	)

	for {
		select {
		case <-done:
			close(inputs)
			return score, comp.Err()
		case inputs <- advent2019.Sign(ballX - paddleX):
			// do nothing
		case x := <-outputs:
			y, z := <-outputs, <-outputs
			if x == -1 && y == 0 {
				score = z
			} else {
				switch tile(z) {
				case tilePaddle:
					paddleX = x
				case tileBall:
					ballX = x
				}
			}
		}
	}
}
