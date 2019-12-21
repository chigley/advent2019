package advent2019

import (
	"container/list"
	"errors"
)

type BFSNode interface {
	IsGoal() bool
	Neighbours() ([]BFSNode, error)
	Key() interface{}
}

func BFS(start BFSNode) (int, error) {
	q := list.New()
	q.PushBack(start)

	parent := map[BFSNode]BFSNode{
		start: nil,
	}

	discovered := map[interface{}]struct{}{
		start: {},
	}

	for e := q.Front(); e != nil; e = e.Next() {
		node := e.Value.(BFSNode)

		if node.IsGoal() {
			var steps int
			for curParent := parent[node]; curParent != nil; curParent = parent[curParent] {
				steps++
			}
			return steps, nil
		}

		neighbours, err := node.Neighbours()
		if err != nil {
			return 0, err
		}
		for _, n := range neighbours {
			if _, ok := discovered[n.Key()]; ok {
				continue
			}
			discovered[n.Key()] = struct{}{}
			parent[n] = node
			q.PushBack(n)
		}
	}

	return 0, errors.New("advent2019: no goal node found")
}
