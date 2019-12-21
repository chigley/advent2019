package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

type Maze struct {
	tiles   map[vector.XY]tile
	portals map[vector.XY]vector.XY
	aa      vector.XY
	zz      vector.XY
}

type tile int

const (
	wall tile = iota
	passage
	portal
)

func NewMaze(r io.Reader) (*Maze, error) {
	lines, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	maze := &Maze{
		tiles:   make(map[vector.XY]tile),
		portals: make(map[vector.XY]vector.XY),
	}

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
			maze.tiles[vector.XY{x, y}] = tile
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
		maze.portals[points[0].entryPoint] = points[1].accessedFrom
		maze.portals[points[1].entryPoint] = points[0].accessedFrom
	}

	return maze, nil
}
