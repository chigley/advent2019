package main

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/chigley/advent2019"
)

var regexpInstr = regexp.MustCompile(`(stack|cut|increment)(?: (-?\d+))?$`)

func ReadOrder(r io.Reader) (Order, error) {
	lines, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	ret := make(Order, len(lines))
	for i, line := range lines {
		matches := regexpInstr.FindStringSubmatch(line)
		if len(matches) != 3 {
			return nil, errors.New("unexpected regexp matches")
		}
		instr, argStr := matches[1], matches[2]

		var arg int
		if argStr != "" {
			arg, err = strconv.Atoi(argStr)
			if err != nil {
				return nil, err
			}
		}

		switch instr {
		case "stack":
			ret[i] = Stack{}
		case "cut":
			ret[i] = Cut{N: arg}
		case "increment":
			ret[i] = Increment{N: arg}
		default:
			return nil, fmt.Errorf("unexpected regexp match %s", instr)
		}
	}
	return ret, nil
}
