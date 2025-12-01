package grid

// Direction represents a cardinal direction.
type Direction int

const (
	// North points upward (negative Y in typical grid coordinates).
	North Direction = iota
	// East points rightward (positive X).
	East
	// South points downward (positive Y in typical grid coordinates).
	South
	// West points leftward (negative X).
	West
)

// AllDirections returns all four cardinal directions.
func AllDirections() []Direction {
	return []Direction{North, East, South, West}
}

// AllDirections8 returns all eight directions including diagonals.
func AllDirections8() []Direction {
	return []Direction{North, East, South, West, NorthEast, SouthEast, SouthWest, NorthWest}
}

const (
	// NorthEast is the diagonal direction between North and East.
	NorthEast Direction = 4 + iota
	// SouthEast is the diagonal direction between South and East.
	SouthEast
	// SouthWest is the diagonal direction between South and West.
	SouthWest
	// NorthWest is the diagonal direction between North and West.
	NorthWest
)

// Offset returns the point offset for a given direction.
func (d Direction) Offset() Point {
	switch d {
	case North:
		return Point{X: 0, Y: -1}
	case East:
		return Point{X: 1, Y: 0}
	case South:
		return Point{X: 0, Y: 1}
	case West:
		return Point{X: -1, Y: 0}
	case NorthEast:
		return Point{X: 1, Y: -1}
	case SouthEast:
		return Point{X: 1, Y: 1}
	case SouthWest:
		return Point{X: -1, Y: 1}
	case NorthWest:
		return Point{X: -1, Y: -1}
	default:
		return Point{X: 0, Y: 0}
	}
}

// Move returns a new point moved in the given direction.
func (d Direction) Move(p Point) Point {
	return p.Add(d.Offset())
}

// TurnRight returns the direction 90 degrees clockwise.
func (d Direction) TurnRight() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		return d
	}
}

// TurnLeft returns the direction 90 degrees counter-clockwise.
func (d Direction) TurnLeft() Direction {
	switch d {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	case West:
		return South
	default:
		return d
	}
}

// Reverse returns the opposite direction.
func (d Direction) Reverse() Direction {
	switch d {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	case NorthEast:
		return SouthWest
	case SouthEast:
		return NorthWest
	case SouthWest:
		return NorthEast
	case NorthWest:
		return SouthEast
	default:
		return d
	}
}

// String returns a string representation of the direction.
func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	case NorthEast:
		return "NorthEast"
	case SouthEast:
		return "SouthEast"
	case SouthWest:
		return "SouthWest"
	case NorthWest:
		return "NorthWest"
	default:
		return "Unknown"
	}
}
