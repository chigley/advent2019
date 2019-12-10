package advent2019

type Point struct {
	X, Y int
}

type Direction struct {
	DX, DY int
}

func (p1 *Point) Direction(p2 Point) Direction {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	gcd := Abs(GCD(dx, dy))

	return Direction{
		DX: dx / gcd,
		DY: dy / gcd,
	}
}

func (p1 *Point) Distance(p2 Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}
