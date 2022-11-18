package common

type LineSegment struct {
	start *Point
	end   *Point
	slope *Point
}

func NewLineSegment(start *Point, end *Point) *LineSegment {
	return &LineSegment{
		start: start,
		end:   end,
	}
}

func (line *LineSegment) Start() *Point {
	return line.start
}

func (line *LineSegment) End() *Point {
	return line.end
}

func (line *LineSegment) Slope() *Point {
	if line.slope == nil {
		x := line.End().X() - line.Start().X()
		y := line.End().Y() - line.Start().Y()
		z := line.End().Z() - line.Start().Z()

		reducesNumbers := Reduce(x, y, z)

		line.slope = New3DPoint(reducesNumbers[0], reducesNumbers[1], reducesNumbers[2])
	}

	return line.slope
}

func (line *LineSegment) IsVertical() bool {
	return line.Start().X() == line.End().X()
}

func (line *LineSegment) IsHorizontal() bool {
	return line.Start().Y() == line.End().Y()
}

func (line *LineSegment) IsStacked() bool {
	return line.Start().Z() == line.End().Z()
}
