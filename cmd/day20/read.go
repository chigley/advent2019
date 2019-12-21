package main

import (
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

type Maze struct {
	tiles   map[vector.XY]tile
	portals map[vector.XY]portalDest
	aa      vector.XY
	zz      vector.XY
}

type tile int

const (
	wall tile = iota
	passage
	portal
)

type portalDest struct {
	pos        vector.XY
	levelDelta int
}

type edges struct {
	minX int
	minY int
	maxX int
	maxY int
}

func NewMaze(r io.Reader) (*Maze, error) {
	lines, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	maze := &Maze{
		tiles:   make(map[vector.XY]tile),
		portals: make(map[vector.XY]portalDest),
	}

	// Keep track of extreme edges, to allow us to distinguish internal portals
	// from external ones
	edges := newEdges()

	for y, line := range lines {
		for x, char := range line {
			var tile tile
			switch char {
			case '#':
				tile = wall
			case '.':
				tile = passage
			default:
				continue
			}
			pos := vector.XY{x, y}
			maze.tiles[pos] = tile
			edges.update(pos)
		}
	}

	type portalLabel string
	type portalData struct {
		entryPoint   vector.XY
		accessedFrom vector.XY
	}
	portalPoints := make(map[portalLabel][]portalData)

	var foundAA, foundZZ bool

	for pos, tile := range maze.tiles {
		if tile != passage {
			continue
		}

		for _, nPos := range pos.Neighbours() {
			if _, ok := maze.tiles[nPos]; ok {
				continue
			}

			// Found a portal
			var char1, char2 uint8
			portalDir, _ := pos.Direction(nPos)
			switch portalDir {
			case vector.South: // north in our coordinate system
				char1 = lines[nPos.Y-1][nPos.X]
				char2 = lines[nPos.Y][nPos.X]
			case vector.East:
				char1 = lines[nPos.Y][nPos.X]
				char2 = lines[nPos.Y][nPos.X+1]
			case vector.North: // south in our coordinate system
				char1 = lines[nPos.Y][nPos.X]
				char2 = lines[nPos.Y+1][nPos.X]
			case vector.West:
				char1 = lines[nPos.Y][nPos.X-1]
				char2 = lines[nPos.Y][nPos.X]
			default:
				return nil, fmt.Errorf("unexpected portal direction %#v", portalDir)
			}

			portalLabel := portalLabel(char1) + portalLabel(char2)
			switch portalLabel {
			case "AA":
				maze.aa = pos
				foundAA = true
			case "ZZ":
				maze.zz = pos
				foundZZ = true
			default:
				portalPoints[portalLabel] = append(portalPoints[portalLabel], portalData{
					entryPoint:   nPos,
					accessedFrom: pos,
				})
				maze.tiles[nPos] = portal
			}
		}
	}

	if !foundAA || !foundZZ {
		return nil, errors.New("failed to find AA, ZZ, or both")
	}

	for label, points := range portalPoints {
		if len(points) != 2 {
			return nil, fmt.Errorf("portal %s has %d points, expected 2", label, len(points))
		}

		maze.portals[points[0].entryPoint] = portalDest{
			pos:        points[1].accessedFrom,
			levelDelta: edges.levelDelta(points[0].accessedFrom),
		}
		maze.portals[points[1].entryPoint] = portalDest{
			pos:        points[0].accessedFrom,
			levelDelta: edges.levelDelta(points[1].accessedFrom),
		}
	}

	return maze, nil
}

func newEdges() *edges {
	return &edges{
		minX: math.MaxInt64,
		minY: math.MaxInt64,
		maxX: math.MinInt64,
		maxY: math.MinInt64,
	}
}

func (e *edges) update(pos vector.XY) {
	e.minY = advent2019.Min(e.minY, pos.Y)
	e.maxY = advent2019.Max(e.maxY, pos.Y)
	e.minX = advent2019.Min(e.minX, pos.X)
	e.maxX = advent2019.Max(e.maxX, pos.X)
}

func (e *edges) levelDelta(pos vector.XY) int {
	if pos.X == e.minX || pos.X == e.maxX || pos.Y == e.minY || pos.Y == e.maxY {
		return -1
	}
	return 1
}
