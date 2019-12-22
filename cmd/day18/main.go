package main

import (
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
	maze        *Maze
	positions   [1]vector.XY
	activeRobot int
	keys        keyMask
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

func (m *Maze) Part1() (int, error) {
	return advent2019.BFS(&node{
		maze:        m,
		positions:   [1]vector.XY{m.entrance},
		activeRobot: -1,
	})
}

func (n *node) IsGoal() bool {
	return n.keys.haveAll(n.maze.numKeys)
}

func (n *node) Neighbours() ([]advent2019.BFSNode, error) {
	if n.activeRobot == -1 {
		ret := make([]advent2019.BFSNode, len(n.positions))
		for i := 0; i < len(n.positions); i++ {
			ret[i] = &node{
				maze:        n.maze,
				positions:   n.positions,
				activeRobot: i,
				keys:        n.keys,
			}
		}
		return ret, nil
	}

	ret := make([]advent2019.BFSNode, 0, 4)
	for _, pos := range n.positions[n.activeRobot].Neighbours() {
		keys := n.keys

		char, ok := n.maze.tiles[pos]
		if !ok || char == '#' {
			continue
		}
		if unicode.IsUpper(char) {
			if !n.keys.haveKey(char) {
				continue
			}
		} else if unicode.IsLower(char) {
			keys = keys.collectKey(char)
		} else if char != '@' && char != '.' {
			return nil, fmt.Errorf("unrecognised character %q", char)
		}

		n.positions[n.activeRobot] = pos

		ret = append(ret, &node{
			maze:      n.maze,
			positions: n.positions,
			keys:      keys,
		})
	}
	return ret, nil
}

func (n *node) Key() interface{} {
	return struct {
		positions [1]vector.XY
		int
		keyMask
	}{
		n.positions,
		n.activeRobot,
		n.keys,
	}
}
