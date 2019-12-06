package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/chigley/advent2019"
)

const com = "COM"

type orbits map[string]string

func main() {
	galaxy, err := readGalaxy(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(galaxy.totalOrbits())
	fmt.Println(galaxy.orbitalTransfers("YOU", "SAN"))
}

func (o orbits) totalOrbits() (total int) {
	for obj := range o {
		total += o.distanceToCOM(obj, 0)
	}
	return
}

func (o orbits) distanceToCOM(obj string, acc int) int {
	if obj == com {
		return acc
	}
	return o.distanceToCOM(o[obj], acc+1)
}

func (o orbits) orbitalTransfers(from, to string) int {
	fromAnc := o.ancestors(from, nil)
	toAnc := o.ancestors(to, nil)

	var i int
	for ; fromAnc[i] == toAnc[i]; i++ {
	}
	return (len(fromAnc) - i) + (len(toAnc) - i)
}

func (o orbits) ancestors(obj string, acc []string) []string {
	if obj == com {
		return acc
	}
	return o.ancestors(o[obj], append([]string{o[obj]}, acc...))
}

func readGalaxy(r io.Reader) (orbits, error) {
	records, err := advent2019.ReadCSV(r, ')')
	if err != nil {
		return nil, err
	}

	galaxy := make(orbits)
	for _, r := range records {
		if len(r) != 2 {
			return nil, fmt.Errorf("invalid input record %v", r)
		}
		if _, ok := galaxy[r[1]]; ok {
			return nil, fmt.Errorf("%s can't orbit two objects", r[1])
		}
		galaxy[r[1]] = r[0]
	}
	return galaxy, nil
}
