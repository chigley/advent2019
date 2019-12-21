package main

import (
	"log"
	"os"
)

func main() {
	maze, err := NewMaze(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	_ = maze
}
