package common

type BoundedPlane struct {
	span      *LineSegment
	locations [][]interface{}
}

func NewBoundedPlane(height int, width int) *BoundedPlane {
	return &BoundedPlane{
		span:      NewLineSegment(New2DPoint(0, 0), New2DPoint(width-1, height-1)),
		locations: initializeLocations(height, width),
	}
}

func NewBoundedPlaneFromValues(values [][]interface{}) *BoundedPlane {
	maxY := len(values) - 1
	maxX := len(values[0]) - 1

	return &BoundedPlane{
		span:      NewLineSegment(New2DPoint(0, 0), New2DPoint(maxX, maxY)),
		locations: values,
	}
}

func NewBoundedPlaneFromPoints(points []*Point) *BoundedPlane {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for _, point := range points {
		minX = MinInt(minX, point.x)
		maxX = MaxInt(maxX, point.x)
		minY = MinInt(minY, point.y)
		maxY = MaxInt(maxY, point.y)
	}

	return &BoundedPlane{
		span:      NewLineSegment(New2DPoint(minX, minY), New2DPoint(maxX, maxY)),
		locations: initializeLocations(maxY-minY+1, maxX-minX+1),
	}
}

func NewBoundedPlaneFromLines(lines []*LineSegment) *BoundedPlane {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for _, line := range lines {
		minX = MinInt(minX, line.start.x, line.end.x)
		maxX = MaxInt(maxX, line.start.x, line.end.x)
		minY = MinInt(minY, line.start.y, line.end.y)
		maxY = MaxInt(maxY, line.start.y, line.end.y)
	}

	return &BoundedPlane{
		span:      NewLineSegment(New2DPoint(minX, minY), New2DPoint(maxX, maxY)),
		locations: initializeLocations(maxY-minY+1, maxX-minX+1),
	}
}

func initializeLocations(height int, width int) [][]interface{} {
	locations := make([][]interface{}, height)

	for row := range locations {

		locations[row] = make([]interface{}, width)
	}

	return locations
}

func (plane *BoundedPlane) Span() *LineSegment {
	return plane.span
}

func (plane *BoundedPlane) GetValueAt(point *Point) interface{} {
	row := point.y - plane.span.start.y
	col := point.x - plane.span.start.x
	return plane.locations[row][col]
}

func (plane *BoundedPlane) SetValueAt(point *Point, value interface{}) {
	row := point.y - plane.span.start.y
	col := point.x - plane.span.start.x
	plane.locations[row][col] = value
}

// Only orthagonally adjacent.
func (plane *BoundedPlane) GetVonNeumannNeighbors(point *Point) []*Point {
	var adjacentPoints []*Point

	if point.x+1 <= plane.span.end.x {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x+1, point.y))
	}

	if point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x, point.y+1))
	}

	if point.x-1 >= plane.span.start.x {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x-1, point.y))
	}

	if point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x, point.y-1))
	}

	return adjacentPoints
}

// Both orthagonally and diagonally adjacent.
func (plane *BoundedPlane) GetMooreNeighbors(point *Point) []*Point {
	var adjacentPoints []*Point

	if point.x+1 <= plane.span.end.x {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x+1, point.y))
	}

	if point.x+1 <= plane.span.end.x && point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x+1, point.y+1))
	}

	if point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x, point.y+1))
	}

	if point.x-1 >= plane.span.start.x && point.y+1 <= plane.span.end.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x-1, point.y+1))
	}

	if point.x-1 >= plane.span.start.x {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x-1, point.y))
	}

	if point.x-1 >= plane.span.start.x && point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x-1, point.y-1))
	}

	if point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x, point.y-1))
	}

	if point.x+1 <= plane.span.end.x && point.y-1 >= plane.span.start.y {
		adjacentPoints = append(adjacentPoints, New2DPoint(point.x+1, point.y-1))
	}

	return adjacentPoints
}

func (plane *BoundedPlane) GetAllPoints() []*Point {
	var points []*Point

	for y := plane.span.start.y; y <= plane.span.end.y; y++ {
		for x := plane.span.start.x; x <= plane.span.end.x; x++ {
			points = append(points, New2DPoint(x, y))
		}
	}

	return points
}
