package grid

import "testing"

func TestDirection_Offset(t *testing.T) {
	tests := []struct {
		name     string
		dir      Direction
		expected Point
	}{
		{"North", North, Point{0, -1}},
		{"East", East, Point{1, 0}},
		{"South", South, Point{0, 1}},
		{"West", West, Point{-1, 0}},
		{"NorthEast", NorthEast, Point{1, -1}},
		{"SouthEast", SouthEast, Point{1, 1}},
		{"SouthWest", SouthWest, Point{-1, 1}},
		{"NorthWest", NorthWest, Point{-1, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.Offset()
			if result != tt.expected {
				t.Errorf("Offset() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDirection_TurnRight(t *testing.T) {
	tests := []struct {
		name     string
		dir      Direction
		expected Direction
	}{
		{"North to East", North, East},
		{"East to South", East, South},
		{"South to West", South, West},
		{"West to North", West, North},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.TurnRight()
			if result != tt.expected {
				t.Errorf("TurnRight() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDirection_TurnLeft(t *testing.T) {
	tests := []struct {
		name     string
		dir      Direction
		expected Direction
	}{
		{"North to West", North, West},
		{"East to North", East, North},
		{"South to East", South, East},
		{"West to South", West, South},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.TurnLeft()
			if result != tt.expected {
				t.Errorf("TurnLeft() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDirection_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		dir      Direction
		expected Direction
	}{
		{"North to South", North, South},
		{"East to West", East, West},
		{"South to North", South, North},
		{"West to East", West, East},
		{"NorthEast to SouthWest", NorthEast, SouthWest},
		{"SouthEast to NorthWest", SouthEast, NorthWest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.Reverse()
			if result != tt.expected {
				t.Errorf("Reverse() = %v, want %v", result, tt.expected)
			}
		})
	}
}

