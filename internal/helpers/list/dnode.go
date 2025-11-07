package list

// dnode represents a double-linked node for use in DoubleLinkedList.
type dnode[T any] struct {
	value T
	prev  *dnode[T]
	next  *dnode[T]
}
