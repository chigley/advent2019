package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	order, err := ReadOrder(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := order.Part1()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
}

func (o Order) Part1() (int, error) {
	deck := factoryDeck(10007)
	deck = o.Shuffle(deck)
	return cardIndex(deck, 2019)
}

func factoryDeck(cards int) []int {
	ret := make([]int, cards)
	for i := 0; i < cards; i++ {
		ret[i] = i
	}
	return ret
}

func cardIndex(deck []int, want int) (int, error) {
	for i, found := range deck {
		if found == want {
			return i, nil
		}
	}
	return 0, errors.New("no solution found")
}
