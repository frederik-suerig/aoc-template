package list

// SortedList represents a linked list that maintains elements in sorted order.
// The comparison function determines the sort order.
type SortedList[T any] struct {
	head     *node[T]
	size     int
	lessThan func(T, T) bool
}

// NewSortedList creates a new sorted list with the provided comparison function.
// lessThan should return true if the first argument is less than the second.
func NewSortedList[T any](lessThan func(T, T) bool) *SortedList[T] {
	return &SortedList[T]{
		lessThan: lessThan,
	}
}

// Insert adds a value to the list in the correct sorted position.
func (l *SortedList[T]) Insert(val T) {
	newNode := &node[T]{value: val}

	if l.head == nil || l.lessThan(val, l.head.value) {
		newNode.next = l.head
		l.head = newNode
		l.size++
		return
	}

	curr := l.head
	for curr.next != nil && l.lessThan(curr.next.value, val) {
		curr = curr.next
	}

	newNode.next = curr.next
	curr.next = newNode
	l.size++
}

// Remove removes the first occurrence of a value from the list.
// Returns true if the value was found and removed.
func (l *SortedList[T]) Remove(val T, equals func(T, T) bool) bool {
	if l.head == nil {
		return false
	}

	if equals(l.head.value, val) {
		l.head = l.head.next
		l.size--
		return true
	}

	curr := l.head
	for curr.next != nil {
		if equals(curr.next.value, val) {
			curr.next = curr.next.next
			l.size--
			return true
		}
		curr = curr.next
	}

	return false
}

// Contains returns true if the list contains the specified value.
func (l *SortedList[T]) Contains(val T, equals func(T, T) bool) bool {
	curr := l.head
	for curr != nil {
		if equals(curr.value, val) {
			return true
		}
		// Early exit if we've passed where the value should be
		if l.lessThan(val, curr.value) {
			return false
		}
		curr = curr.next
	}
	return false
}

// Get returns the value at the specified index and true if the index is valid.
// Returns zero value and false if the index is out of bounds.
func (l *SortedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, false
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.value, true
}

// Size returns the number of elements in the list.
func (l *SortedList[T]) Size() int {
	return l.size
}

// IsEmpty returns true if the list is empty.
func (l *SortedList[T]) IsEmpty() bool {
	return l.size == 0
}

// ToSlice returns all values in the list as a slice (in sorted order).
func (l *SortedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	curr := l.head
	for curr != nil {
		result = append(result, curr.value)
		curr = curr.next
	}
	return result
}
