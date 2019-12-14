package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
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

	// The cloning is not strictly necessary here: the part 2 answer will be the
	// same if we start computing it 1,000 steps in.
	fmt.Println(Part1(moons.Clone(), 1000))
	fmt.Println(Part2(moons))
}

func Part1(ms Moons, steps int) int {
	ms.stepN(steps)
	return ms.energy()
}

func Part2(ms Moons) int {
	return advent2019.LCM(
		ms.X.phase(),
		ms.Y.phase(),
		ms.Z.phase(),
	)
}

func (ms *Moons) Clone() (ret Moons) {
	ret.X = make(Axis, len(ms.X))
	copy(ret.X, ms.X)

	ret.Y = make(Axis, len(ms.Y))
	copy(ret.Y, ms.Y)

	ret.Z = make(Axis, len(ms.Z))
	copy(ret.Z, ms.Z)

	return
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

func (a Axis) phase() (phase int) {
	orig := make(Axis, len(a))
	copy(orig, a)

	for {
		a.step()
		phase++
		if reflect.DeepEqual(orig, a) {
			return
		}
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
