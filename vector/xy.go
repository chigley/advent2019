package vector

import (
	"github.com/chigley/advent2019"
)

type XY struct {
	X, Y int
}

func (p1 *XY) Direction(p2 XY) (XY, int) {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	gcd := advent2019.Abs(advent2019.GCD(dx, dy))

	return XY{
		X: dx / gcd,
		Y: dy / gcd,
	}, gcd
}

func (p1 *XY) Distance(p2 XY) int {
	return advent2019.Abs(p1.X-p2.X) + advent2019.Abs(p1.Y-p2.Y)
}
