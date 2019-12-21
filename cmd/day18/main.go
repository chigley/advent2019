package main

import (
	"container/list"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

type Maze struct {
	tiles    map[vector.XY]rune
	entrance vector.XY
	numKeys  uint
}

type node struct {
	pos  vector.XY
	keys keyMask
}

func main() {
	maze, err := NewMaze(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := maze.Part1()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(part1)
}

func NewMaze(r io.Reader) (*Maze, error) {
	lines, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	m := &Maze{
		tiles: make(map[vector.XY]rune),
	}
	for y, line := range lines {
		for x, char := range line {
			pos := vector.XY{x, y}
			m.tiles[pos] = char
			if char == '@' {
				m.entrance = pos
			}
			if unicode.IsLower(char) {
				m.numKeys++
			}
		}
	}
	return m, nil
}

func (m Maze) Part1() (int, error) {
	startNode := node{
		pos: m.entrance,
	}

	q := list.New()
	q.PushBack(startNode)

	discovered := map[node]struct{}{
		startNode: {},
	}
	steps := make(map[vector.XY]int)

	for e := q.Front(); e != nil; e = e.Next() {
		node := e.Value.(node)

		char, ok := m.tiles[node.pos]
		if !ok || char == '#' {
			continue
		}
		if unicode.IsUpper(char) {
			if !node.keys.haveKey(char) {
				continue
			}
		} else if unicode.IsLower(char) {
			node.keys.collectKey(char)
			if node.keys.haveAll(m.numKeys) {
				return steps[node.pos], nil
			}
		} else if char != '@' && char != '.' {
			return 0, fmt.Errorf("unrecognised character %q", char)
		}

		nodeSteps := steps[node.pos]
		for _, neighbour := range node.neighbours() {
			if _, ok := discovered[neighbour]; !ok {
				discovered[neighbour] = struct{}{}
				steps[neighbour.pos] = nodeSteps + 1
				q.PushBack(neighbour)
			}
		}
	}

	return 0, errors.New("no solution found")
}

func (n *node) neighbours() []node {
	ret := make([]node, 0, 4)
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			if x == 0 && y == 0 {
				// same point
				continue
			}

			if advent2019.Abs(x) == 1 && advent2019.Abs(y) == 1 {
				// diagonal
				continue
			}

			pos := vector.XY{n.pos.X + x, n.pos.Y + y}
			ret = append(ret, node{
				pos:  pos,
				keys: n.keys,
			})
		}
	}
	return ret
}
