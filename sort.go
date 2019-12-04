package advent2019

import (
	"sort"
)

type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p runeSlice) Sort() { sort.Sort(p) }

func StringIsSorted(s string) bool {
	return sort.IsSorted(runeSlice(s))
}
