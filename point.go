package advent2019

type Point struct {
	X, Y int
}

func (p1 *Point) Distance(p2 Point) int {
	return Abs(p1.X-p2.X) + Abs(p1.Y-p2.Y)
}
