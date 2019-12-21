package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

type node struct {
	pos  vector.XY
	maze *Maze
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

	fmt.Println(part1)
}

func (m *Maze) Part1() (int, error) {
	return advent2019.BFS(&node{
		pos:  m.aa,
		maze: m,
	})
}

func (n *node) IsGoal() bool {
	return n.pos == n.maze.zz
}

func (n *node) Neighbours() ([]advent2019.BFSNode, error) {
	ret := make([]advent2019.BFSNode, 0, 4)
	for _, dir := range vector.Dirs {
		pos := n.pos.Add(dir)
		tile, ok := n.maze.tiles[pos]
		if !ok {
			continue
		}
		switch tile {
		case wall:
			continue
		case passage:
			ret = append(ret, &node{
				pos:  pos,
				maze: n.maze,
			})
		case portal:
			portalDest, ok := n.maze.portals[pos]
			if !ok {
				return nil, fmt.Errorf("found a portal tile but don't know where it goes")
			}
			ret = append(ret, &node{
				pos:  portalDest,
				maze: n.maze,
			})
		default:
			return nil, fmt.Errorf("unexpected tile type %d", tile)
		}
	}
	return ret, nil
}

func (n *node) Key() interface{} {
	return n.pos
}
