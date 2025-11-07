package list

// node represents a single-linked node for use in FIFO, Stack, and LinkedList.
type node[T any] struct {
	value T
	next  *node[T]
}
