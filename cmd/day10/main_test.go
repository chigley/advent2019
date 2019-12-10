package main_test

import (
	"os"
	"testing"

	"github.com/chigley/advent2019"
	day10 "github.com/chigley/advent2019/cmd/day10"
	"github.com/stretchr/testify/assert"
)

var part1Tests = []struct {
	input  []string
	output int
}{
	{
		[]string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##",
		},
		8,
	},
	{
		[]string{
			"......#.#.",
			"#..#.#....",
			"..#######.",
			".#.#.###..",
			".#..#.....",
			"..#....#.#",
			"#..#....#.",
			".##.#..###",
			"##...#..#.",
			".#....####",
		},
		33,
	},
	{
		[]string{
			"#.#...#.#.",
			".###....#.",
			".#....#...",
			"##.#.#.#.#",
			"....#.#.#.",
			".##..###.#",
			"..#...##..",
			"..##....##",
			"......#...",
			".####.###.",
		},
		35,
	},
	{
		[]string{
			".#..#..###",
			"####.###.#",
			"....###.#.",
			"..###.##.#",
			"##.##.#.#.",
			"....###..#",
			"..#.#..#.#",
			"#..#.#.###",
			".##...##.#",
			".....#.#..",
		},
		41,
	},
	{
		[]string{
			".#..##.###...#######",
			"##.############..##.",
			".#.######.########.#",
			".###.#######.####.#.",
			"#####.##.#.##.###.##",
			"..#####..#.#########",
			"####################",
			"#.####....###.#.#.##",
			"##.#################",
			"#####.##.###..####..",
			"..######..##.#######",
			"####.##.####...##..#",
			".#####..#.######.###",
			"##...#.##########...",
			"#.##########.#######",
			".####.#.###.###.#.##",
			"....##.##.###..#####",
			".#.#.###########.###",
			"#.#.#.#####.####.###",
			"###.##.####.##.#..##",
		},
		210,
	},
}

var part2Tests = []struct {
	input []string
	hit   map[int]advent2019.Point
}{
	{
		[]string{
			".#..##.###...#######",
			"##.############..##.",
			".#.######.########.#",
			".###.#######.####.#.",
			"#####.##.#.##.###.##",
			"..#####..#.#########",
			"####################",
			"#.####....###.#.#.##",
			"##.#################",
			"#####.##.###..####..",
			"..######..##.#######",
			"####.##.####...##..#",
			".#####..#.######.###",
			"##...#.##########...",
			"#.##########.#######",
			".####.#.###.###.#.##",
			"....##.##.###..#####",
			".#.#.###########.###",
			"#.#.#.#####.####.###",
			"###.##.####.##.#..##",
		},
		map[int]advent2019.Point{
			0:   {11, 12},
			1:   {12, 1},
			2:   {12, 2},
			9:   {12, 8},
			19:  {16, 0},
			49:  {16, 9},
			99:  {10, 16},
			198: {9, 6},
			199: {8, 2},
			200: {10, 9},
			298: {11, 1},
		},
	},
}

func TestDay10(t *testing.T) {
	input, err := os.Open("testdata/input")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	lines, err := advent2019.ReadStrings(input)
	if err != nil {
		t.Fatal(err)
	}

	asteroids := day10.Asteroids(lines)

	laserPos, part1 := day10.Part1(asteroids)
	assert.Equal(t, 280, part1)

	part2, err := day10.Part2(laserPos, asteroids)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 706, 100*part2.X+part2.Y)
}

func TestPart1(t *testing.T) {
	for _, tt := range part1Tests {
		_, output := day10.Part1(day10.Asteroids(tt.input))
		assert.Equal(t, tt.output, output)
	}
}

func TestPart2(t *testing.T) {
	for _, tt := range part2Tests {
		asteroids := day10.Asteroids(tt.input)
		laserPos, _ := day10.Part1(asteroids)
		hits := day10.HitOrder(laserPos, asteroids)
		for i, expected := range tt.hit {
			assert.Equal(t, expected, hits[i])
		}
	}
}
