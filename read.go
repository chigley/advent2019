package advent2019

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Range struct {
	From int
	To   int
}

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

func ReadCSV(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	return reader.ReadAll()
}

func ReadRanges(r io.Reader) ([]Range, error) {
	reader := csv.NewReader(r)
	reader.Comma = '-'
	reader.FieldsPerRecord = 2

	var ranges []Range
	for {
		record, err := reader.Read()
		if err == io.EOF {
			return ranges, nil
		}
		if err != nil {
			return nil, err
		}

		if len(record) != 2 {
			return nil, fmt.Errorf("advent2019: invalid range %s", record)
		}
		from, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("advent2019: invalid range %s", record)
		}
		to, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("advent2019: invalid range %s", record)
		}

		if from > to {
			return nil, fmt.Errorf("advent2019: invalid range %s", record)
		}

		ranges = append(ranges, Range{From: from, To: to})
	}
}
