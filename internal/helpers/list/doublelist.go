package list

// DoubleLinkedList represents a double-linked list with bidirectional traversal.
type DoubleLinkedList[T any] struct {
	head *dnode[T]
	tail *dnode[T]
	size int
}

// NewDoubleLinkedList creates a new empty double linked list.
func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

// Append adds a value to the end of the list.
func (l *DoubleLinkedList[T]) Append(val T) {
	newNode := &dnode[T]{value: val}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}

	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
	l.size++
}

// Prepend adds a value to the beginning of the list.
func (l *DoubleLinkedList[T]) Prepend(val T) {
	newNode := &dnode[T]{value: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
	l.size++
}

// Remove removes the first occurrence of a value from the list.
// Returns true if the value was found and removed.
func (l *DoubleLinkedList[T]) Remove(val T, equals func(T, T) bool) bool {
	curr := l.head
	for curr != nil {
		if equals(curr.value, val) {
			if curr.prev != nil {
				curr.prev.next = curr.next
			} else {
				l.head = curr.next
			}

			if curr.next != nil {
				curr.next.prev = curr.prev
			} else {
				l.tail = curr.prev
			}

			l.size--
			return true
		}
		curr = curr.next
	}

	return false
}

// Get returns the value at the specified index and true if the index is valid.
// Returns zero value and false if the index is out of bounds.
// Uses forward traversal for indices in the first half, backward for the second half.
func (l *DoubleLinkedList[T]) Get(index int) (T, bool) {
	if index < 0 || index >= l.size {
		var zero T
		return zero, false
	}

	// Optimize: traverse from the end if index is in the second half
	if index > l.size/2 {
		curr := l.tail
		for i := l.size - 1; i > index; i-- {
			curr = curr.prev
		}
		return curr.value, true
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.value, true
}

// Set updates the value at the specified index.
// Returns true if the index is valid.
func (l *DoubleLinkedList[T]) Set(index int, val T) bool {
	if index < 0 || index >= l.size {
		return false
	}

	// Optimize: traverse from the end if index is in the second half
	if index > l.size/2 {
		curr := l.tail
		for i := l.size - 1; i > index; i-- {
			curr = curr.prev
		}
		curr.value = val
		return true
	}

	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	curr.value = val
	return true
}

// Insert inserts a value at the specified index.
// Returns true if the index is valid.
func (l *DoubleLinkedList[T]) Insert(index int, val T) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.Prepend(val)
		return true
	}

	if index == l.size {
		l.Append(val)
		return true
	}

	// Find the node at the insertion point
	var target *dnode[T]
	if index > l.size/2 {
		target = l.tail
		for i := l.size - 1; i >= index; i-- {
			target = target.prev
		}
	} else {
		target = l.head
		for i := 0; i < index-1; i++ {
			target = target.next
		}
	}

	newNode := &dnode[T]{
		value: val,
		prev:  target,
		next:  target.next,
	}

	if target.next != nil {
		target.next.prev = newNode
	}
	target.next = newNode

	if newNode.next == nil {
		l.tail = newNode
	}

	l.size++
	return true
}

// Size returns the number of elements in the list.
func (l *DoubleLinkedList[T]) Size() int {
	return l.size
}

// IsEmpty returns true if the list is empty.
func (l *DoubleLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// ToSlice returns all values in the list as a slice (from head to tail).
func (l *DoubleLinkedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	curr := l.head
	for curr != nil {
		result = append(result, curr.value)
		curr = curr.next
	}
	return result
}

// ToSliceReverse returns all values in the list as a slice (from tail to head).
func (l *DoubleLinkedList[T]) ToSliceReverse() []T {
	result := make([]T, 0, l.size)
	curr := l.tail
	for curr != nil {
		result = append(result, curr.value)
		curr = curr.prev
	}
	return result
}
