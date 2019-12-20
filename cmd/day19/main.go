package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
	"github.com/chigley/advent2019/vector"
)

type Drone struct {
	cache map[vector.XY]bool
	comp  *intcode.Computer
	maxX  int
	maxY  int
}

func main() {
	program, err := advent2019.ReadIntsLine(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	drone := NewDrone(program)

	part1, err := drone.Part1()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(drone)

	fmt.Println(part1)
}

func NewDrone(program []int) *Drone {
	return &Drone{
		cache: make(map[vector.XY]bool),
		comp:  intcode.New(program),
	}
}

func (d *Drone) Part1() (int, error) {
	var count int
	for x := 0; x < 50; x++ {
		for y := 0; y < 50; y++ {
			isPulled, err := d.readPos(vector.XY{x, y})
			if err != nil {
				return 0, err
			}
			if isPulled {
				count++
			}
		}
	}
	return count, nil
}

func (d *Drone) String() string {
	var b strings.Builder
	for x := 0; x <= d.maxX; x++ {
		for y := 0; y <= d.maxY; y++ {
			isPulled, ok := d.cache[vector.XY{x, y}]
			if ok {
				switch isPulled {
				case true:
					b.WriteRune('#')
				default:
					b.WriteRune('.')
				}
			} else {
				b.WriteRune('?')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (d *Drone) readPos(pos vector.XY) (bool, error) {
	if res, ok := d.cache[pos]; ok {
		return res, nil
	}

	out, err := d.comp.Run([]int{pos.X, pos.Y})
	if err != nil {
		return false, err
	}
	if len(out) != 1 {
		return false, fmt.Errorf("expected 1 output, got %d", len(out))
	}

	d.maxX = advent2019.Max(d.maxX, pos.X)
	d.maxY = advent2019.Max(d.maxY, pos.Y)

	res := out[0] == 1
	d.cache[pos] = res
	return res, nil
}
