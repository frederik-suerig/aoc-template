package grid

// Point represents a 2D coordinate point.
type Point struct {
	X, Y int
}

// NewPoint creates a new point with the given coordinates.
func NewPoint(x, y int) Point {
	return Point{X: x, Y: y}
}

// Add returns a new point that is the sum of this point and another.
func (p Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

// Subtract returns a new point that is the difference of this point and another.
func (p Point) Subtract(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

// Multiply returns a new point scaled by the given factor.
func (p Point) Multiply(factor int) Point {
	return Point{X: p.X * factor, Y: p.Y * factor}
}

// ManhattanDistance returns the Manhattan distance between two points.
func (p Point) ManhattanDistance(other Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

// ChebyshevDistance returns the Chebyshev (chessboard) distance between two points.
func (p Point) ChebyshevDistance(other Point) int {
	return max(abs(p.X-other.X), abs(p.Y-other.Y))
}

// EuclideanDistanceSquared returns the squared Euclidean distance (avoids floating point).
func (p Point) EuclideanDistanceSquared(other Point) int {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return dx*dx + dy*dy
}

// abs returns the absolute value of an integer.
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
