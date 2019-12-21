package main

import (
	"unicode"
)

type keyMask int

func (m keyMask) collectKey(k rune) keyMask {
	return m | 1<<keyIndex(k)
}

func (m keyMask) haveKey(k rune) bool {
	mask := 1 << keyIndex(k)
	return int(m)&mask == mask
}

func (m keyMask) haveAll(target uint) bool {
	return m == (1<<target)-1
}

func keyIndex(k rune) uint {
	return uint(unicode.ToLower(k)) - uint('a')
}
