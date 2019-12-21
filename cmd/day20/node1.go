package main

import (
	"fmt"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

type node1 struct {
	pos  vector.XY
	maze *Maze
}

func (n *node1) IsGoal() bool {
	return n.pos == n.maze.zz
}

func (n *node1) Neighbours() ([]advent2019.BFSNode, error) {
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
			ret = append(ret, &node1{
				pos:  pos,
				maze: n.maze,
			})
		case portal:
			portalDest, ok := n.maze.portals[pos]
			if !ok {
				return nil, fmt.Errorf("found a portal tile but don't know where it goes")
			}
			ret = append(ret, &node1{
				pos:  portalDest.pos,
				maze: n.maze,
			})
		default:
			return nil, fmt.Errorf("unexpected tile type %d", tile)
		}
	}
	return ret, nil
}

func (n *node1) Key() interface{} {
	return n.pos
}
