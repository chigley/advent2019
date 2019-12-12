package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
	"github.com/chigley/advent2019/vector"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

type colour int

const (
	black colour = iota
	white
)

type pointColours map[vector.XY]colour

type robot struct {
	computer   *intcode.Computer
	grid       pointColours
	paintedSet map[vector.XY]struct{}

	pos vector.XY
	dir direction
}

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
	fmt.Print(part2)
}

func Part1(program []int) (int, error) {
	bot := robot{
		computer:   intcode.New(program),
		grid:       make(pointColours),
		paintedSet: make(map[vector.XY]struct{}),
	}
	if err := bot.draw(); err != nil {
		return 0, err
	}
	return len(bot.paintedSet), nil
}

func Part2(program []int) (string, error) {
	bot := robot{
		computer:   intcode.New(program),
		grid:       make(pointColours),
		paintedSet: make(map[vector.XY]struct{}),
	}
	bot.grid[vector.XY{X: 0, Y: 0}] = white
	if err := bot.draw(); err != nil {
		return "", err
	}
	return bot.grid.String(), nil
}

func (r *robot) draw() error {
	done := make(chan struct{})
	inputs := make(chan int)
	outputs := r.computer.RunInteractive(inputs, func() {
		close(done)
	})

	for {
		select {
		case <-done:
			close(inputs)
			return r.computer.Err()
		case inputs <- int(r.grid[r.pos]):
			newCol, turnDir := <-outputs, <-outputs

			r.paintedSet[r.pos] = struct{}{}

			r.grid[r.pos] = colour(newCol)
			r.dir = r.dir.rotate(turnDir)

			dx, dy := r.dir.delta()
			r.pos.X += dx
			r.pos.Y += dy
		}
	}
}

func (d direction) delta() (int, int) {
	switch d {
	case north:
		return 0, 1
	case east:
		return 1, 0
	case south:
		return 0, -1
	default:
		return -1, 0
	}
}

func (d direction) rotate(input int) direction {
	if input == 0 {
		return (d + 3) % 4
	}
	return (d + 1) % 4
}

func (g pointColours) String() string {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	for pos := range g {
		minX = advent2019.Min(minX, pos.X)
		minY = advent2019.Min(minY, pos.Y)
		maxX = advent2019.Max(maxX, pos.X)
		maxY = advent2019.Max(maxY, pos.Y)
	}

	var b strings.Builder

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			col := g[vector.XY{X: x, Y: y}]
			if col == black {
				b.WriteString(".")
			} else {
				b.WriteString("#")
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}
