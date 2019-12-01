package advent2019

import (
	"bufio"
	"io"
	"strconv"
)

func ReadInts(r io.Reader) ([]int, error) {
	var vals []int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}

	return vals, scanner.Err()
}
