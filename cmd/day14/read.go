package main

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/chigley/advent2019"
)

var regexpReaction = regexp.MustCompile(`^(\d+ [A-Z]+(?:, \d+ [A-Z]+)*) => (\d+ [A-Z]+)$`)

func ReadReactions(r io.Reader) (Reactions, error) {
	input, err := advent2019.ReadStrings(r)
	if err != nil {
		return nil, err
	}

	ret := make(Reactions, len(input))
	for _, line := range input {
		matches := regexpReaction.FindStringSubmatch(line)

		var quantities []ReactionInput
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
		ret[out.Chem] = Reaction{
			Ins:  ins,
			OutN: out.N,
		}
	}
	return ret, nil
}

func parseQuantity(s string) (ReactionInput, error) {
	var (
		chem Chemical
		n    int
	)
	if _, err := fmt.Sscanf(s, "%d %s", &n, &chem); err != nil {
		return ReactionInput{}, err
	}
	return ReactionInput{
		Chem: chem,
		N:    n,
	}, nil
}
