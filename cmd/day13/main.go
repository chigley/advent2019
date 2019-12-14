package main

import (
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
	fmt.Println(part1)
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
