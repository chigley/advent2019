package main

import (
	"fmt"
	"io"
	"regexp"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

var regexpMoon = regexp.MustCompile(`^<x=(-?\d+), y=(-?\d+), z=(-?\d+)>$`)

func ReadMoons(r io.Reader) (Moons, error) {
	input, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	ret := make([]Moon, len(input))
	for i, in := range input {
		strMatches := regexpMoon.FindStringSubmatch(in)
		if len(strMatches) != 4 {
			return nil, fmt.Errorf("failed to parse moon '%s'", in)
		}

		matches, err := advent2019.AtoiSlice(strMatches[1:])
		if err != nil {
			return nil, err
		}

		ret[i] = Moon{
			Pos: vector.XYZ{
				X: matches[0],
				Y: matches[1],
				Z: matches[2],
			},
		}
	}
	return ret, nil
}
