package main

import (
	"fmt"
	"log"
	"os"
)

type reactions map[chemical]reaction

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

	fmt.Printf("%#v", reactions)
}
