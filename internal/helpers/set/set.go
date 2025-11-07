package set

// Set represents a collection of unique values, similar to Python's set.
type Set[T comparable] struct {
	data map[T]struct{}
	size int
}

// NewSet creates a new empty set.
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}),
		size: 0,
	}
}

// NewSetWithCapacity creates a new empty set with the specified initial capacity.
func NewSetWithCapacity[T comparable](capacity int) *Set[T] {
	return &Set[T]{
		data: make(map[T]struct{}, capacity),
		size: 0,
	}
}

// Add adds a value to the set if it's not already present.
// Returns true if the value was added (was not already in the set).
func (s *Set[T]) Add(val T) bool {
	if _, exists := s.data[val]; exists {
		return false
	}
	s.data[val] = struct{}{}
	s.size++
	return true
}

// Remove removes a value from the set if it's present.
// Returns true if the value was removed (was in the set).
func (s *Set[T]) Remove(val T) bool {
	if _, exists := s.data[val]; !exists {
		return false
	}
	delete(s.data, val)
	s.size--
	return true
}

// Contains returns true if the value is in the set.
func (s *Set[T]) Contains(val T) bool {
	_, exists := s.data[val]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set[T]) Size() int {
	return s.size
}

// IsEmpty returns true if the set is empty.
func (s *Set[T]) IsEmpty() bool {
	return s.size == 0
}

// ToSlice returns all values in the set as a slice.
// The order is not guaranteed (map iteration order in Go is random).
func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0, s.size)
	for val := range s.data {
		result = append(result, val)
	}
	return result
}

// Union creates a new set containing all elements from both sets.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSetWithCapacity[T](s.size + other.size)
	for val := range s.data {
		result.Add(val)
	}
	for val := range other.data {
		result.Add(val)
	}
	return result
}

// Intersection creates a new set containing only elements present in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	// Iterate over the smaller set for efficiency
	if s.size <= other.size {
		for val := range s.data {
			if other.Contains(val) {
				result.Add(val)
			}
		}
	} else {
		for val := range other.data {
			if s.Contains(val) {
				result.Add(val)
			}
		}
	}
	return result
}

// Difference creates a new set containing elements in this set but not in the other.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSetWithCapacity[T](s.size)
	for val := range s.data {
		if !other.Contains(val) {
			result.Add(val)
		}
	}
	return result
}

// IsSubset returns true if this set is a subset of the other set.
func (s *Set[T]) IsSubset(other *Set[T]) bool {
	if s.size > other.size {
		return false
	}
	for val := range s.data {
		if !other.Contains(val) {
			return false
		}
	}
	return true
}
