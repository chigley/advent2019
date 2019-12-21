package main

import (
	"fmt"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

func (n *node2) IsGoal() bool {
	return n.pos == n.maze.zz && n.level == 0
}

func (n *node2) Neighbours() ([]advent2019.BFSNode, error) {
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
			ret = append(ret, &node2{
				level: n.level,
				node1: node1{
					pos:  pos,
					maze: n.maze,
				},
			})
		case portal:
			portalDest, ok := n.maze.portals[pos]
			if !ok {
				return nil, fmt.Errorf("found a portal tile but don't know where it goes")
			}

			if portalDest.levelDelta == -1 && n.level == 0 {
				// Outer portals don't work on level 0.
				continue
			}

			ret = append(ret, &node2{
				level: n.level + portalDest.levelDelta,
				node1: node1{
					pos:  portalDest.pos,
					maze: n.maze,
				},
			})
		default:
			return nil, fmt.Errorf("unexpected tile type %d", tile)
		}
	}
	return ret, nil
}

func (n *node2) Key() interface{} {
	return struct {
		vector.XY
		int
	}{
		n.pos,
		n.level,
	}
}
