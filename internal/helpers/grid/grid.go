package grid

import (
	"fmt"
)

// Grid represents a 2D grid of values.
type Grid[T any] struct {
	data   [][]T
	width  int
	height int
}

// NewGrid creates a new grid with the specified dimensions and initial value.
func NewGrid[T any](width, height int, initial T) *Grid[T] {
	data := make([][]T, height)
	for y := range data {
		data[y] = make([]T, width)
		for x := range data[y] {
			data[y][x] = initial
		}
	}
	return &Grid[T]{
		data:   data,
		width:  width,
		height: height,
	}
}

// NewGridFromData creates a new grid from existing 2D slice data.
// The data is not copied, so modifications to the grid will affect the original slice.
func NewGridFromData[T any](data [][]T) (*Grid[T], error) {
	if len(data) == 0 {
		return &Grid[T]{width: 0, height: 0}, nil
	}

	width := len(data[0])
	for i, row := range data {
		if len(row) != width {
			return nil, fmt.Errorf("row %d has length %d, expected %d", i, len(row), width)
		}
	}

	return &Grid[T]{
		data:   data,
		width:  width,
		height: len(data),
	}, nil
}

// ParseGrid parses a grid from string lines, converting each character using the converter function.
func ParseGrid[T any](lines []string, converter func(rune) T) (*Grid[T], error) {
	if len(lines) == 0 {
		return &Grid[T]{width: 0, height: 0}, nil
	}

	height := len(lines)
	width := len(lines[0])

	// Validate all lines have the same width
	for i, line := range lines {
		if len(line) != width {
			return nil, fmt.Errorf("line %d has length %d, expected %d", i, len(line), width)
		}
	}

	data := make([][]T, height)
	for y, line := range lines {
		data[y] = make([]T, width)
		for x, r := range line {
			data[y][x] = converter(r)
		}
	}

	return &Grid[T]{
		data:   data,
		width:  width,
		height: height,
	}, nil
}

// ParseStringGrid creates a grid where each cell is a rune from the input strings.
func ParseStringGrid(lines []string) (*Grid[rune], error) {
	return ParseGrid(lines, func(r rune) rune { return r })
}

// Width returns the width of the grid.
func (g *Grid[T]) Width() int {
	return g.width
}

// Height returns the height of the grid.
func (g *Grid[T]) Height() int {
	return g.height
}

// Get returns the value at the given point and true if the point is within bounds.
// Returns zero value and false if out of bounds.
func (g *Grid[T]) Get(p Point) (T, bool) {
	return g.GetXY(p.X, p.Y)
}

// GetXY returns the value at the given coordinates and true if the coordinates are within bounds.
// Returns zero value and false if out of bounds.
func (g *Grid[T]) GetXY(x, y int) (T, bool) {
	if !g.InBoundsXY(x, y) {
		var zero T
		return zero, false
	}
	return g.data[y][x], true
}

// Set sets the value at the given point.
// Returns true if the point is within bounds.
func (g *Grid[T]) Set(p Point, value T) bool {
	return g.SetXY(p.X, p.Y, value)
}

// SetXY sets the value at the given coordinates.
// Returns true if the coordinates are within bounds.
func (g *Grid[T]) SetXY(x, y int, value T) bool {
	if !g.InBoundsXY(x, y) {
		return false
	}
	g.data[y][x] = value
	return true
}

// InBounds returns true if the point is within the grid bounds.
func (g *Grid[T]) InBounds(p Point) bool {
	return g.InBoundsXY(p.X, p.Y)
}

// InBoundsXY returns true if the coordinates are within the grid bounds.
func (g *Grid[T]) InBoundsXY(x, y int) bool {
	return x >= 0 && x < g.width && y >= 0 && y < g.height
}

// Neighbors returns all valid neighbors of a point in the given directions.
func (g *Grid[T]) Neighbors(p Point, directions []Direction) []Point {
	var neighbors []Point
	for _, dir := range directions {
		neighbor := dir.Move(p)
		if g.InBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

// Neighbors4 returns the four cardinal neighbors (North, East, South, West) that are within bounds.
func (g *Grid[T]) Neighbors4(p Point) []Point {
	return g.Neighbors(p, AllDirections())
}

// Neighbors8 returns all eight neighbors (including diagonals) that are within bounds.
func (g *Grid[T]) Neighbors8(p Point) []Point {
	return g.Neighbors(p, AllDirections8())
}

// ForEach calls the function for each cell in the grid.
// The function receives the point, value, and returns true to continue, false to stop.
func (g *Grid[T]) ForEach(fn func(Point, T) bool) {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			p := Point{X: x, Y: y}
			if !fn(p, g.data[y][x]) {
				return
			}
		}
	}
}

// Find returns the first point where the predicate returns true, and true if found.
func (g *Grid[T]) Find(predicate func(Point, T) bool) (Point, bool) {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			p := Point{X: x, Y: y}
			if predicate(p, g.data[y][x]) {
				return p, true
			}
		}
	}
	return Point{}, false
}

// FindAll returns all points where the predicate returns true.
func (g *Grid[T]) FindAll(predicate func(Point, T) bool) []Point {
	var results []Point
	g.ForEach(func(p Point, v T) bool {
		if predicate(p, v) {
			results = append(results, p)
		}
		return true
	})
	return results
}

// Copy creates a deep copy of the grid.
func (g *Grid[T]) Copy() *Grid[T] {
	data := make([][]T, g.height)
	for y := range g.data {
		data[y] = make([]T, g.width)
		copy(data[y], g.data[y])
	}
	return &Grid[T]{
		data:   data,
		width:  g.width,
		height: g.height,
	}
}

// Row returns a copy of the row at the given y coordinate.
func (g *Grid[T]) Row(y int) []T {
	if y < 0 || y >= g.height {
		return nil
	}
	row := make([]T, g.width)
	copy(row, g.data[y])
	return row
}

// Col returns a copy of the column at the given x coordinate.
func (g *Grid[T]) Col(x int) []T {
	if x < 0 || x >= g.width {
		return nil
	}
	col := make([]T, g.height)
	for y := 0; y < g.height; y++ {
		col[y] = g.data[y][x]
	}
	return col
}
