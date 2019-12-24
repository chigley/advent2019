package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

const size = 5

type grid map[vector.XY]tile

type tile int

const (
	tileEmpty tile = iota
	tileBug
)

func main() {
	grid, err := ReadGrid(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(grid))
}

func Part1(g grid) uint64 {
	seen := map[uint64]struct{}{
		g.rating(): {},
	}

	for {
		g = g.step()
		rating := g.rating()
		if _, ok := seen[rating]; ok {
			return rating
		}
		seen[rating] = struct{}{}
	}
}

func (g grid) step() grid {
	ret := make(grid)
	for pos, tile := range g {
		bugs := g.adjacentBugs(pos)
		if tile == tileBug && bugs != 1 {
			ret[pos] = tileEmpty
		} else if tile == tileEmpty && (bugs == 1 || bugs == 2) {
			ret[pos] = tileBug
		} else {
			ret[pos] = tile
		}
	}
	return ret
}

func (g grid) adjacentBugs(pos vector.XY) int {
	var ret int
	for _, nPos := range pos.Neighbours() {
		if g[nPos] == tileBug {
			ret++
		}
	}
	return ret
}

func (g grid) rating() uint64 {
	var (
		rating uint64
		base   uint
	)
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			if g[vector.XY{x, y}] == tileBug {
				rating |= 1 << base
			}
			base++
		}
	}
	return rating
}

func (g grid) String() string {
	var b strings.Builder
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			tile := g[vector.XY{x, y}]
			switch tile {
			case tileEmpty:
				b.WriteRune('.')
			case tileBug:
				b.WriteRune('#')
			default:
				b.WriteRune('?')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func ReadGrid(r io.Reader) (grid, error) {
	lines, err := advent2019.ReadStrings(r)
	if err != nil {
		log.Fatal(err)
	}

	grid := make(grid)
	for y, line := range lines {
		for x, char := range line {
			pos := vector.XY{x, y}
			switch char {
			case '.':
				grid[pos] = tileEmpty
			case '#':
				grid[pos] = tileBug
			default:
				return nil, fmt.Errorf("unrecognised input character %q", char)
			}
		}
	}
	return grid, nil
}
