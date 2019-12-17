package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/intcode"
)

type tile int

const (
	tileOpen tile = iota
	tileScaffold
	tileRobotOnScaffold
	tileRobotOffScaffold
)

const (
	routineMain = "A,B,A,C,A,B,C,A,B,C"
	routineA    = "R,12,R,4,R,10,R,12"
	routineB    = "R,6,L,8,R,10"
	routineC    = "L,8,R,4,R,4,R,6"
)

func main() {
	program, err := advent2019.ReadIntsLine(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	view, err := ReadView(program)
	if err != nil {
		log.Fatal(err)
	}

	part1, err := Part1(view)
	if err != nil {
		log.Fatal(err)
	}

	part2, err := Part2(program)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func Part1(view string) (int, error) {
	tiles, err := genMap(view)
	if err != nil {
		return 0, err
	}

	height := len(tiles)
	if height < 3 {
		return 0, errors.New("need at least three rows")
	}
	width := len(tiles[0])
	if width < 3 {
		return 0, errors.New("need at least three columns")
	}

	var sum int
	for top := 1; top < height-1; top++ {
	tile:
		for left := 1; left < width-1; left++ {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if advent2019.Abs(i) == 1 && advent2019.Abs(j) == 1 {
						// diagonal: not interesting
						continue
					}
					if !tiles[top+i][left+j].isScaffold() {
						continue tile
					}
				}
			}
			sum += top * left
		}
	}
	return sum, nil
}

func Part2(program []int) (int, error) {
	if len(program) < 1 {
		return 0, errors.New("program too short to overwrite first index")
	}
	program[0] = 2

	inputStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n",
		routineMain,
		routineA,
		routineB,
		routineC,
		"n",
	)

	inputs := make([]int, len(inputStr))
	for i, char := range inputStr {
		inputs[i] = int(char)
	}

	outputs, err := intcode.New(program).Run(inputs)
	if err != nil {
		return 0, err
	}
	if len(outputs) == 0 {
		return 0, errors.New("no output")
	}
	return outputs[len(outputs)-1], nil
}

func ReadView(program []int) (string, error) {
	outputs, err := intcode.New(program).Run(nil)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	for _, out := range outputs {
		b.WriteRune(rune(out))
	}
	return b.String(), nil
}

func genMap(view string) ([][]tile, error) {
	lines := strings.Split(strings.TrimSpace(view), "\n")
	height := len(lines)
	if height == 0 {
		return nil, errors.New("need at least one row")
	}
	width := len(lines[0])

	ret := make([][]tile, height)
	for i := height - 1; i >= 0; i-- {
		ret[i] = make([]tile, width)
		for j := 0; j < width; j++ {
			switch lines[i][j] {
			case '.':
				ret[i][j] = tileOpen
			case '#':
				ret[i][j] = tileScaffold
			case '^', 'V', '<', '>':
				ret[i][j] = tileRobotOnScaffold
			case 'X':
				ret[i][j] = tileRobotOffScaffold
			default:
				return nil, fmt.Errorf("unrecognised character %v", rune(lines[i][j]))
			}
		}
	}
	return ret, nil
}

func (t tile) isScaffold() bool {
	return t == tileScaffold || t == tileRobotOnScaffold
}
