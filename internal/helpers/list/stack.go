package list

// Stack represents a last-in-first-out stack.
type Stack[T any] struct {
	top  *node[T]
	size int
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(val T) {
	newNode := &node[T]{
		value: val,
		next:  s.top,
	}
	s.top = newNode
	s.size++
}

// Pop removes and returns the value at the top of the stack.
// Returns the value and true if the stack was not empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.top == nil {
		var zero T
		return zero, false
	}

	val := s.top.value
	s.top = s.top.next
	s.size--
	return val, true
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of elements in the stack.
func (s *Stack[T]) Size() int {
	return s.size
}

// ToSlice returns all values in the stack as a slice (from top to bottom).
func (s *Stack[T]) ToSlice() []T {
	result := make([]T, 0, s.size)
	curr := s.top
	for curr != nil {
		result = append(result, curr.value)
		curr = curr.next
	}
	return result
}
