package advent2019

import (
	"bufio"
	"encoding/csv"
	"errors"
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

func ReadIntsLine(r io.Reader) ([]int, error) {
	reader := csv.NewReader(r)

	strs, err := reader.Read()
	if err == io.EOF {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if _, err := reader.Read(); err != io.EOF {
		return nil, errors.New("advent2019: unexpected extra input")
	}

	vals := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		val, err := strconv.Atoi(strs[i])
		if err != nil {
			return nil, err
		}
		vals[i] = val
	}

	return vals, nil
}
