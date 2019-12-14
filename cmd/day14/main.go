package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/chigley/advent2019"
)

type (
	Reactions map[Chemical]Reaction
	store     map[Chemical]int
)

type Chemical string

type Reaction struct {
	Ins  []ReactionInput
	OutN int
}

type ReactionInput struct {
	Chem Chemical
	N    int
}

func main() {
	reactions, err := ReadReactions(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reactions.Part1(1))
	fmt.Println(reactions.Part2(1000000000000))
}

func (r Reactions) Part1(fuel int) int {
	have := make(store)
	want := ReactionInput{"FUEL", fuel}

	ore, _ := r.oreReqd(have, want)
	return ore
}

func (r Reactions) Part2(ore int) int {
	return sort.Search(ore+1, func(i int) bool {
		return r.Part1(i) > ore
	}) - 1
}

func (r Reactions) oreReqd(have store, want ReactionInput) (int, store) {
	if want.Chem == "ORE" {
		return want.N, have
	}
	if want.N == 0 {
		return 0, have
	}

	useFromStock := advent2019.Min(want.N, have[want.Chem])
	want.N -= useFromStock
	have[want.Chem] -= useFromStock

	numReacts := (want.N + r[want.Chem].OutN - 1) / r[want.Chem].OutN
	spare := numReacts*r[want.Chem].OutN - want.N

	have[want.Chem] += spare

	var ore int
	for _, want := range r[want.Chem].Ins {
		want.N *= numReacts

		var additionalOre int
		additionalOre, have = r.oreReqd(have, want)
		ore += additionalOre
	}

	return ore, have
}
