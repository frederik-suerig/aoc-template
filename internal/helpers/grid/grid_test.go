package grid

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid[int](3, 2, 0)
	if g.Width() != 3 {
		t.Errorf("Width() = %d, want 3", g.Width())
	}
	if g.Height() != 2 {
		t.Errorf("Height() = %d, want 2", g.Height())
	}

	// Check all values are initialized
	g.ForEach(func(p Point, v int) bool {
		if v != 0 {
			t.Errorf("Expected 0 at %v, got %d", p, v)
		}
		return true
	})
}

func TestParseStringGrid(t *testing.T) {
	lines := []string{
		"ABC",
		"DEF",
		"GHI",
	}

	g, err := ParseStringGrid(lines)
	if err != nil {
		t.Fatalf("ParseStringGrid() error = %v", err)
	}

	if g.Width() != 3 {
		t.Errorf("Width() = %d, want 3", g.Width())
	}
	if g.Height() != 3 {
		t.Errorf("Height() = %d, want 3", g.Height())
	}

	// Check values
	if val, ok := g.GetXY(0, 0); !ok || val != 'A' {
		t.Errorf("GetXY(0, 0) = %c, %v, want 'A', true", val, ok)
	}
	if val, ok := g.GetXY(1, 1); !ok || val != 'E' {
		t.Errorf("GetXY(1, 1) = %c, %v, want 'E', true", val, ok)
	}
}

func TestGrid_GetSet(t *testing.T) {
	g := NewGrid[int](3, 3, 0)

	p := Point{X: 1, Y: 1}
	if !g.Set(p, 42) {
		t.Error("Set() = false, want true")
	}

	val, ok := g.Get(p)
	if !ok {
		t.Error("Get() = _, false, want _, true")
	}
	if val != 42 {
		t.Errorf("Get() = %d, want 42", val)
	}
}

func TestGrid_InBounds(t *testing.T) {
	g := NewGrid[int](3, 3, 0)

	tests := []struct {
		name     string
		p        Point
		expected bool
	}{
		{"inside", Point{1, 1}, true},
		{"top-left", Point{0, 0}, true},
		{"bottom-right", Point{2, 2}, true},
		{"negative x", Point{-1, 1}, false},
		{"negative y", Point{1, -1}, false},
		{"too large x", Point{3, 1}, false},
		{"too large y", Point{1, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := g.InBounds(tt.p)
			if result != tt.expected {
				t.Errorf("InBounds() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGrid_Neighbors4(t *testing.T) {
	g := NewGrid[int](3, 3, 0)

	// Center point should have 4 neighbors
	p := Point{X: 1, Y: 1}
	neighbors := g.Neighbors4(p)
	if len(neighbors) != 4 {
		t.Errorf("Neighbors4() returned %d neighbors, want 4", len(neighbors))
	}

	// Corner point should have 2 neighbors
	corner := Point{X: 0, Y: 0}
	cornerNeighbors := g.Neighbors4(corner)
	if len(cornerNeighbors) != 2 {
		t.Errorf("Neighbors4() at corner returned %d neighbors, want 2", len(cornerNeighbors))
	}
}

func TestGrid_Neighbors8(t *testing.T) {
	g := NewGrid[int](3, 3, 0)

	// Center point should have 8 neighbors
	p := Point{X: 1, Y: 1}
	neighbors := g.Neighbors8(p)
	if len(neighbors) != 8 {
		t.Errorf("Neighbors8() returned %d neighbors, want 8", len(neighbors))
	}

	// Corner point should have 3 neighbors
	corner := Point{X: 0, Y: 0}
	cornerNeighbors := g.Neighbors8(corner)
	if len(cornerNeighbors) != 3 {
		t.Errorf("Neighbors8() at corner returned %d neighbors, want 3", len(cornerNeighbors))
	}
}

func TestGrid_Find(t *testing.T) {
	g := NewGrid[int](3, 3, 0)
	g.SetXY(1, 1, 42)

	p, found := g.Find(func(_ Point, v int) bool {
		return v == 42
	})

	if !found {
		t.Error("Find() = _, false, want _, true")
	}
	if p.X != 1 || p.Y != 1 {
		t.Errorf("Find() = %v, want {1, 1}", p)
	}
}

func TestGrid_FindAll(t *testing.T) {
	g := NewGrid[int](3, 3, 0)
	g.SetXY(0, 0, 42)
	g.SetXY(2, 2, 42)

	points := g.FindAll(func(_ Point, v int) bool {
		return v == 42
	})

	if len(points) != 2 {
		t.Errorf("FindAll() returned %d points, want 2", len(points))
	}
}
