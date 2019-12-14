package main_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	day14 "github.com/chigley/advent2019/cmd/day14"
)

var tests = []struct {
	input day14.Reactions
	part1 int
	part2 int
}{
	{
		input: day14.Reactions{
			"DCFZ":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 165}}, OutN: 6},
			"FUEL":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "XJWVT", N: 44}, {Chem: "KHKGT", N: 5}, {Chem: "QDVJ", N: 1}, {Chem: "NZVS", N: 29}, {Chem: "GPVTF", N: 9}, {Chem: "HKGWZ", N: 48}}, OutN: 1},
			"GPVTF": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 165}}, OutN: 2},
			"HKGWZ": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 177}}, OutN: 5},
			"KHKGT": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "DCFZ", N: 3}, {Chem: "NZVS", N: 7}, {Chem: "HKGWZ", N: 5}, {Chem: "PSHF", N: 10}}, OutN: 8},
			"NZVS":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 157}}, OutN: 5}, "PSHF": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 179}}, OutN: 7},
			"QDVJ":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "HKGWZ", N: 12}, {Chem: "GPVTF", N: 1}, {Chem: "PSHF", N: 8}}, OutN: 9},
			"XJWVT": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "DCFZ", N: 7}, {Chem: "PSHF", N: 7}}, OutN: 2},
		},
		part1: 13312,
		part2: 82892753,
	},
	{
		input: day14.Reactions{
			"CXFTF": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "NVRVD", N: 1}}, OutN: 8},
			"FUEL":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "STKFG", N: 53}, {Chem: "MNCFX", N: 6}, {Chem: "VJHF", N: 46}, {Chem: "HVMC", N: 81}, {Chem: "CXFTF", N: 68}, {Chem: "GNMV", N: 25}}, OutN: 1},
			"FWMGM": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "VJHF", N: 22}, {Chem: "MNCFX", N: 37}}, OutN: 5},
			"GNMV":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "VJHF", N: 5}, {Chem: "MNCFX", N: 7}, {Chem: "VPVL", N: 9}, {Chem: "CXFTF", N: 37}}, OutN: 6},
			"HVMC":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "MNCFX", N: 5}, {Chem: "RFSQX", N: 7}, {Chem: "FWMGM", N: 2}, {Chem: "VPVL", N: 2}, {Chem: "CXFTF", N: 19}}, OutN: 3}, "JNWZP": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 144}}, OutN: 7},
			"MNCFX": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 145}}, OutN: 6},
			"NVRVD": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 139}}, OutN: 4},
			"RFSQX": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "VJHF", N: 1}, {Chem: "MNCFX", N: 6}}, OutN: 4},
			"STKFG": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "VPVL", N: 2}, {Chem: "FWMGM", N: 7}, {Chem: "CXFTF", N: 2}, {Chem: "MNCFX", N: 11}}, OutN: 1},
			"VJHF":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 176}}, OutN: 6},
			"VPVL":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "NVRVD", N: 17}, {Chem: "JNWZP", N: 3}}, OutN: 8},
		},
		part1: 180697,
		part2: 5586022,
	},
	{
		input: day14.Reactions{
			"BHXH":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 114}}, OutN: 4},
			"BMBT":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "VRPVC", N: 14}}, OutN: 6},
			"CNZTR": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 171}}, OutN: 8},
			"FHTLT": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "WPTQ", N: 6}, {Chem: "BMBT", N: 2}, {Chem: "ZLQW", N: 8}, {Chem: "KTJDG", N: 18}, {Chem: "XMNCP", N: 1}, {Chem: "MZWV", N: 6}, {Chem: "RJRHP", N: 1}}, OutN: 6},
			"FUEL":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "BHXH", N: 6}, {Chem: "KTJDG", N: 18}, {Chem: "WPTQ", N: 12}, {Chem: "PLWSL", N: 7}, {Chem: "FHTLT", N: 31}, {Chem: "ZDVW", N: 37}}, OutN: 1},
			"KTJDG": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 189}}, OutN: 9}, "LTCX": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "BHXH", N: 5}, {Chem: "VRPVC", N: 4}}, OutN: 5},
			"MZWV":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "BHXH", N: 3}, {Chem: "VRPVC", N: 2}}, OutN: 7},
			"PLWSL": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ZLQW", N: 7}, {Chem: "BMBT", N: 3}, {Chem: "XCVML", N: 9}, {Chem: "XMNCP", N: 26}, {Chem: "WPTQ", N: 1}, {Chem: "MZWV", N: 2}, {Chem: "RJRHP", N: 1}}, OutN: 4},
			"RJRHP": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "XCVML", N: 7}}, OutN: 6}, "VRPVC": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "ORE", N: 121}}, OutN: 7},
			"WPTQ":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "BMBT", N: 5}}, OutN: 4},
			"XCVML": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "KTJDG", N: 15}, {Chem: "BHXH", N: 12}}, OutN: 5}, "XDBXC": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "VRPVC", N: 12}, {Chem: "CNZTR", N: 27}}, OutN: 2},
			"XMNCP": day14.Reaction{Ins: []day14.ReactionInput{{Chem: "MZWV", N: 1}, {Chem: "XDBXC", N: 17}, {Chem: "XCVML", N: 3}}, OutN: 2},
			"ZDVW":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "WPTQ", N: 13}, {Chem: "LTCX", N: 10}, {Chem: "RJRHP", N: 3}, {Chem: "XMNCP", N: 14}, {Chem: "MZWV", N: 2}, {Chem: "ZLQW", N: 1}}, OutN: 1},
			"ZLQW":  day14.Reaction{Ins: []day14.ReactionInput{{Chem: "XDBXC", N: 15}, {Chem: "LTCX", N: 2}, {Chem: "VRPVC", N: 1}}, OutN: 6}},
		part1: 2210736,
		part2: 460664,
	},
}

func TestDay14(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	reactions, err := day14.ReadReactions(input)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 301997, reactions.Part1(1))
	assert.Equal(t, 6216589, reactions.Part2(1000000000000))

	for _, tt := range tests {
		assert.Equal(t, tt.part1, tt.input.Part1(1))
		assert.Equal(t, tt.part2, tt.input.Part2(1000000000000))
	}
}
