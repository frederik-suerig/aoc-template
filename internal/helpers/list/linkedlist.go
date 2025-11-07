package list

// LinkedList represents a basic single-linked list.
type LinkedList[T any] struct {
	head *node[T]
	size int
}

// NewLinkedList creates a new empty linked list.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append adds a value to the end of the list.
func (l *LinkedList[T]) Append(val T) {
	newNode := &node[T]{value: val}

	if l.head == nil {
		l.head = newNode
		l.size++
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
	l.size++
}

// Prepend adds a value to the beginning of the list.
func (l *LinkedList[T]) Prepend(val T) {
	newNode := &node[T]{
		value: val,
		next:  l.head,
	}
	l.head = newNode
	l.size++
}

// Remove removes the first occurrence of a value from the list.
// Returns true if the value was found and removed.
func (l *LinkedList[T]) Remove(val T, equals func(T, T) bool) bool {
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

// Get returns the value at the specified index and true if the index is valid.
// Returns zero value and false if the index is out of bounds.
func (l *LinkedList[T]) Get(index int) (T, bool) {
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

// Set updates the value at the specified index.
// Returns true if the index is valid.
func (l *LinkedList[T]) Set(index int, val T) bool {
	if index < 0 || index >= l.size {
		return false
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
func (l *LinkedList[T]) Insert(index int, val T) bool {
	if index < 0 || index > l.size {
		return false
	}

	if index == 0 {
		l.Prepend(val)
		return true
	}

	curr := l.head
	for i := 0; i < index-1; i++ {
		curr = curr.next
	}

	newNode := &node[T]{
		value: val,
		next:  curr.next,
	}
	curr.next = newNode
	l.size++
	return true
}

// Size returns the number of elements in the list.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// IsEmpty returns true if the list is empty.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// ToSlice returns all values in the list as a slice.
func (l *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	curr := l.head
	for curr != nil {
		result = append(result, curr.value)
		curr = curr.next
	}
	return result
}
