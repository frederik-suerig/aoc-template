package list

// FIFO represents a first-in-first-out queue.
type FIFO[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

// Push adds a value to the end of the queue.
func (q *FIFO[T]) Push(val T) {
	newNode := &node[T]{value: val}

	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
		q.size++
		return
	}

	q.tail.next = newNode
	q.tail = newNode
	q.size++
}

// Pop removes and returns the value at the front of the queue.
// Returns the value and true if the queue was not empty.
func (q *FIFO[T]) Pop() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}

	val := q.head.value
	q.head = q.head.next
	q.size--

	if q.head == nil {
		q.tail = nil
	}

	return val, true
}

// IsEmpty returns true if the queue is empty.
func (q *FIFO[T]) IsEmpty() bool {
	return q.head == nil
}

// Size returns the number of elements in the queue.
func (q *FIFO[T]) Size() int {
	return q.size
}

// ToSlice returns all values in the queue as a slice (from front to back).
func (q *FIFO[T]) ToSlice() []T {
	result := make([]T, 0, q.size)
	curr := q.head
	for curr != nil {
		result = append(result, curr.value)
		curr = curr.next
	}
	return result
}
