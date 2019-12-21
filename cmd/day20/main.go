package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chigley/advent2019"
)

type node2 struct {
	level int
	node1
}

func main() {
	maze, err := NewMaze(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := maze.Part1()
	if err != nil {
		log.Fatal(err)
	}

	part2, err := maze.Part2()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func (m *Maze) Part1() (int, error) {
	return advent2019.BFS(&node1{
		pos:  m.aa,
		maze: m,
	})
}

func (m *Maze) Part2() (int, error) {
	return advent2019.BFS(&node2{
		node1: node1{
			pos:  m.aa,
			maze: m,
		},
	})
}
