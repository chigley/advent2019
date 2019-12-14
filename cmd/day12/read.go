package main

import (
	"fmt"
	"io"
	"regexp"

	"github.com/chigley/advent2019"
)

var regexpMoon = regexp.MustCompile(`^<x=(-?\d+), y=(-?\d+), z=(-?\d+)>$`)

func ReadMoons(r io.Reader) (Moons, error) {
	input, err := advent2019.ReadStrings(r)
	if err != nil {
		return Moons{}, err
	}

	ret := Moons{
		X: make(Axis, len(input)),
		Y: make(Axis, len(input)),
		Z: make(Axis, len(input)),
	}
	for i, in := range input {
		strMatches := regexpMoon.FindStringSubmatch(in)
		if len(strMatches) != 4 {
			return Moons{}, fmt.Errorf("failed to parse moon '%s'", in)
		}

		matches, err := advent2019.AtoiSlice(strMatches[1:])
		if err != nil {
			return Moons{}, err
		}

		ret.X[i].Pos = matches[0]
		ret.Y[i].Pos = matches[1]
		ret.Z[i].Pos = matches[2]
	}
	return ret, nil
}
