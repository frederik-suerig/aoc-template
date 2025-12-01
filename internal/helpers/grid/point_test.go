package grid

import "testing"

func TestPoint_Add(t *testing.T) {
	tests := []struct {
		name     string
		p1       Point
		p2       Point
		expected Point
	}{
		{"zero points", Point{0, 0}, Point{0, 0}, Point{0, 0}},
		{"positive", Point{1, 2}, Point{3, 4}, Point{4, 6}},
		{"negative", Point{-1, -2}, Point{-3, -4}, Point{-4, -6}},
		{"mixed", Point{1, -2}, Point{-3, 4}, Point{-2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p1.Add(tt.p2)
			if result != tt.expected {
				t.Errorf("Add() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_ManhattanDistance(t *testing.T) {
	tests := []struct {
		name     string
		p1       Point
		p2       Point
		expected int
	}{
		{"same point", Point{0, 0}, Point{0, 0}, 0},
		{"horizontal", Point{0, 0}, Point{5, 0}, 5},
		{"vertical", Point{0, 0}, Point{0, 5}, 5},
		{"diagonal", Point{0, 0}, Point{3, 4}, 7},
		{"negative", Point{-1, -2}, Point{2, 3}, 8},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p1.ManhattanDistance(tt.p2)
			if result != tt.expected {
				t.Errorf("ManhattanDistance() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestPoint_ChebyshevDistance(t *testing.T) {
	tests := []struct {
		name     string
		p1       Point
		p2       Point
		expected int
	}{
		{"same point", Point{0, 0}, Point{0, 0}, 0},
		{"horizontal", Point{0, 0}, Point{5, 0}, 5},
		{"vertical", Point{0, 0}, Point{0, 5}, 5},
		{"diagonal", Point{0, 0}, Point{3, 4}, 4},
		{"negative", Point{-1, -2}, Point{2, 3}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p1.ChebyshevDistance(tt.p2)
			if result != tt.expected {
				t.Errorf("ChebyshevDistance() = %d, want %d", result, tt.expected)
			}
		})
	}
}

