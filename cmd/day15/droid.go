package main

import (
	"container/list"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
)

type tile int

const (
	unexplored tile = iota
	wall
	traversable
	oxygen
)

type Droid struct {
	inputs, outputs chan int

	pos   pos
	tiles map[pos]tile
}

func NewDroid(program []int) *Droid {
	inputs := make(chan int)
	return &Droid{
		inputs:  inputs,
		outputs: intcode.New(program).RunInteractive(inputs, nil),
	}
}

func (d *Droid) bfs(wantOxygen bool) (int, error) {
	d.tiles = make(map[pos]tile)

	q := list.New()
	q.PushBack(list.New())

	longestPath := -1

	for e := q.Front(); e != nil; e = e.Next() {
		path := e.Value.(*list.List)
		history := list.New()

		// Traverse the path we popped from the queue
		var madeProgress bool
		var dist int
		for dirElem := path.Front(); dirElem != nil; dirElem = dirElem.Next() {
			dist++
			dir := dirElem.Value.(direction)
			startPos := d.pos
			if err := d.move(dir); err != nil {
				return 0, err
			}
			if d.pos != startPos {
				if d.tiles[d.pos] == oxygen && wantOxygen {
					return dist, nil
				}
				if dirElem == path.Back() {
					madeProgress = true
				}
				longestPath = advent2019.Max(longestPath, dist)
				history.PushBack(dir)
			}
		}

		// If we reached a new square (or if we just started) queue all new
		// squares reachable from here
		if path.Len() == 0 || madeProgress {
			for _, dir := range directions {
				candidatePos, err := d.pos.next(dir)
				if err != nil {
					return 0, err
				}
				if _, ok := d.tiles[candidatePos]; !ok {
					candidatePath := list.New()
					candidatePath.PushBackList(path)
					candidatePath.PushBack(dir)

					q.PushBack(candidatePath)
				}
			}
		}

		// Traverse back to our origin point
		for dir := history.Back(); dir != nil; dir = dir.Prev() {
			oppositeMove, err := dir.Value.(direction).opposite()
			if err != nil {
				return 0, err
			}
			if err := d.move(oppositeMove); err != nil {
				return 0, err
			}
		}
	}

	if wantOxygen {
		return 0, errors.New("not found")
	}
	return longestPath, nil
}

func (d *Droid) move(dir direction) error {
	d.inputs <- int(dir)

	attemptedPos, err := d.pos.next(dir)
	if err != nil {
		return err
	}

	status := <-d.outputs
	switch status {
	case 0:
		d.tiles[attemptedPos] = wall
		return nil
	case 1, 2:
		d.pos = attemptedPos
		if status == 1 {
			d.tiles[attemptedPos] = traversable
		} else {
			d.tiles[attemptedPos] = oxygen
		}
		return nil
	default:
		return fmt.Errorf("unexpected status code %d", status)
	}
}

func (d *Droid) String() string {
	minX, minY := math.MaxInt64, math.MaxInt64
	maxX, maxY := math.MinInt64, math.MinInt64
	for pos := range d.tiles {
		minX = advent2019.Min(minX, pos.X)
		minY = advent2019.Min(minY, pos.Y)
		maxX = advent2019.Max(maxX, pos.X)
		maxY = advent2019.Max(maxY, pos.Y)
	}

	var b strings.Builder
	b.WriteString("\n")
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			pos := pos{X: x, Y: y}
			if pos == d.pos {
				b.WriteRune('D')
			} else {
				tile := d.tiles[pos]
				switch tile {
				case unexplored:
					b.WriteRune(' ')
				case wall:
					b.WriteRune('#')
				case traversable:
					b.WriteRune('.')
				case oxygen:
					b.WriteRune('O')
				default:
					b.WriteRune('?')
				}
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
