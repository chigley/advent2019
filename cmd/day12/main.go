package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

type Moons struct {
	X, Y, Z Axis
}

type Axis []struct {
	Pos, Vel int
}

func main() {
	moons, err := ReadMoons(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(moons, 1000))
}

func Part1(ms Moons, steps int) int {
	ms.stepN(steps)
	return ms.energy()
}

func (ms *Moons) step() {
	ms.X.step()
	ms.Y.step()
	ms.Z.step()
}

func (ms *Moons) stepN(n int) {
	for i := 0; i < n; i++ {
		ms.step()
	}
}

func (ms *Moons) energy() (ret int) {
	for i := 0; i < len(ms.X); i++ {
		pos := vector.XYZ{X: ms.X[i].Pos, Y: ms.Y[i].Pos, Z: ms.Z[i].Pos}
		vel := vector.XYZ{X: ms.X[i].Vel, Y: ms.Y[i].Vel, Z: ms.Z[i].Vel}
		ret += pos.SumAbs() * vel.SumAbs()
	}
	return
}

func (a Axis) step() {
	for i, m1 := range a {
		for j, m2 := range a {
			if i >= j {
				continue
			}

			delta := advent2019.Sign(m2.Pos - m1.Pos)
			a[i].Vel += delta
			a[j].Vel -= delta
		}
	}

	for i, m := range a {
		a[i].Pos += m.Vel
	}
}

func (ms *Moons) String() string {
	var b strings.Builder
	for i := 0; i < len(ms.X); i++ {
		fmt.Fprintf(&b, "pos=<x=%d, y=%d, z=%d>, vel=<x=%d, y=%d, z=%d>\n",
			ms.X[i].Pos, ms.Y[i].Pos, ms.Z[i].Pos,
			ms.X[i].Vel, ms.Y[i].Vel, ms.Z[i].Vel,
		)
	}
	return b.String()
}
