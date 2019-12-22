package intcode

import "strings"

func ToASCII(input string) []int {
	ret := make([]int, len(input))
	for i, char := range input {
		ret[i] = int(char)
	}
	return ret
}

func FromASCII(output []int) string {
	var b strings.Builder
	for _, out := range output {
		b.WriteRune(rune(out))
	}
	return b.String()
}
