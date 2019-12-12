package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chigley/advent2019"
	"github.com/chigley/advent2019/vector"
)

type Moons []Moon

type Moon struct {
	Pos, Vel vector.XYZ
}

func main() {
	moons, err := ReadMoons(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(Part1(moons, 1000))
}

func Part1(ms Moons, steps int) int {
	ms = ms.stepN(steps)
	return ms.energy()
}

func (ms Moons) stepN(n int) Moons {
	for i := 0; i < n; i++ {
		ms = ms.step()
	}
	return ms
}

func (ms Moons) step() Moons {
	ret := make(Moons, len(ms))
	copy(ret, ms)

	for i, m1 := range ms {
		for j, m2 := range ms {
			if i >= j {
				continue
			}

			x := advent2019.Sign(m2.Pos.X - m1.Pos.X)
			ret[i].Vel.X += x
			ret[j].Vel.X -= x

			y := advent2019.Sign(m2.Pos.Y - m1.Pos.Y)
			ret[i].Vel.Y += y
			ret[j].Vel.Y -= y

			z := advent2019.Sign(m2.Pos.Z - m1.Pos.Z)
			ret[i].Vel.Z += z
			ret[j].Vel.Z -= z
		}
	}

	for i, m := range ret {
		ret[i].Pos = m.Pos.Add(m.Vel)
	}

	return ret
}

func (ms Moons) energy() (ret int) {
	for _, m := range ms {
		ret += m.energy()
	}
	return
}

func (m Moon) energy() int {
	return m.Pos.SumAbs() * m.Vel.SumAbs()
}

func (ms Moons) String() string {
	var b strings.Builder
	for _, m := range ms {
		fmt.Fprintf(&b, "pos=<%v>, vel=<%v>\n", m.Pos, m.Vel)
	}
	return b.String()
}
