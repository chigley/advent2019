package vector

import "github.com/chigley/advent2019"

type XYZ struct {
	X, Y, Z int
}

func (p1 *XYZ) Add(p2 XYZ) XYZ {
	return XYZ{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
		Z: p1.Z + p2.Z,
	}
}

func (p *XYZ) SumAbs() int {
	return advent2019.Abs(p.X) + advent2019.Abs(p.Y) + advent2019.Abs(p.Z)
}
