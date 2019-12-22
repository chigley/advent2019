package main

import "github.com/chigley/advent2019"

type Order []Shuffler

type Shuffler interface {
	Shuffle(d []int) []int
}

type Stack struct{}

type Cut struct {
	N int
}

type Increment struct {
	N int
}

func (o Order) Shuffle(deck []int) []int {
	for _, instr := range o {
		deck = instr.Shuffle(deck)
	}
	return deck
}

func (s Stack) Shuffle(deck []int) []int {
	for i := len(deck)/2 - 1; i >= 0; i-- {
		opp := len(deck) - 1 - i
		deck[i], deck[opp] = deck[opp], deck[i]
	}
	return deck
}

func (c Cut) Shuffle(deck []int) []int {
	if c.N == 0 {
		return deck
	}
	if c.N > 0 {
		return append(deck[c.N:], deck[:c.N]...)
	}
	c.N = advent2019.Abs(c.N)
	return append(deck[len(deck)-c.N:], deck[:len(deck)-c.N]...)
}

func (i Increment) Shuffle(deck []int) []int {
	ret := make([]int, len(deck))
	for j, card := range deck {
		ret[(j*i.N)%len(deck)] = card
	}
	return ret
}
