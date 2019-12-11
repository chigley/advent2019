package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
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

type robot struct {
	computer   *intcode.Computer
	grid       map[advent2019.Point]colour
	paintedSet map[advent2019.Point]struct{}

	pos advent2019.Point
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

	fmt.Println(part1)
}

func Part1(program []int) (int, error) {
	bot := robot{
		computer:   intcode.New(program),
		grid:       make(map[advent2019.Point]colour),
		paintedSet: make(map[advent2019.Point]struct{}),
	}
	if err := bot.draw(); err != nil {
		return 0, err
	}
	return len(bot.paintedSet), nil
}

func (r *robot) draw() error {
	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan struct{})

	inputs := make(chan int)
	outputs := r.computer.RunInteractive(inputs, func() {
		close(done)
	})

	go func() {
		for {
			select {
			case <-done:
				wg.Done()
				return
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
	}()

	wg.Wait()
	close(inputs)

	return r.computer.Err()
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
