package main

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/chigley/advent2019"
)

var regexpReaction = regexp.MustCompile(`^(\d+ [A-Z]+(?:, \d+ [A-Z]+)*) => (\d+ [A-Z]+)$`)

func readReactions(r io.Reader) (reactions, error) {
	input, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	ret := make(reactions, len(input))
	for _, line := range input {
		matches := regexpReaction.FindStringSubmatch(line)

		var quantities []reactionInput
		for _, inStmt := range matches[1:] {
			for _, in := range strings.Split(inStmt, ", ") {
				q, err := parseQuantity(in)
				if err != nil {
					return nil, err
				}
				quantities = append(quantities, q)
			}
		}

		ins, out := quantities[:len(quantities)-1], quantities[len(quantities)-1]
		ret[out.chem] = reaction{
			ins:  ins,
			outN: out.n,
		}
	}
	return ret, nil
}

func parseQuantity(s string) (reactionInput, error) {
	var (
		chem chemical
		n    int
	)
	if _, err := fmt.Sscanf(s, "%d %s", &n, &chem); err != nil {
		return reactionInput{}, err
	}
	return reactionInput{
		chem: chem,
		n:    n,
	}, nil
}
