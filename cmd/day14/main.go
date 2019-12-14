package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/chigley/advent2019"
)

type (
	reactions map[chemical]reaction
	store     map[chemical]int
)

type chemical string

type reaction struct {
	ins  []reactionInput
	outN int
}

type reactionInput struct {
	chem chemical
	n    int
}

func main() {
	reactions, err := readReactions(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reactions.Part1(1))
	fmt.Println(reactions.Part2(1000000000000))
}

func (r reactions) Part1(fuel int) int {
	have := make(store)
	want := reactionInput{"FUEL", fuel}

	ore, _ := r.oreReqd(have, want)
	return ore
}

func (r reactions) Part2(ore int) int {
	return sort.Search(ore+1, func(i int) bool {
		return r.Part1(i) > ore
	}) - 1
}

func (r reactions) oreReqd(have store, want reactionInput) (int, store) {
	if want.chem == "ORE" {
		return want.n, have
	}
	if want.n == 0 {
		return 0, have
	}

	useFromStock := advent2019.Min(want.n, have[want.chem])
	want.n -= useFromStock
	have[want.chem] -= useFromStock

	numReacts := (want.n + r[want.chem].outN - 1) / r[want.chem].outN
	spare := numReacts*r[want.chem].outN - want.n

	have[want.chem] += spare

	var ore int
	for _, want := range r[want.chem].ins {
		want.n *= numReacts

		var additionalOre int
		additionalOre, have = r.oreReqd(have, want)
		ore += additionalOre
	}

	return ore, have
}
