package circularlinkedlist

// Node represents a single node in a circular linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a circular singly linked list.
type LinkedList[T any] struct {
	tail *Node[T]
	size int
}

// New creates and returns a new empty LinkedList.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append adds an element after the tail. The new node becomes the tail.
func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value}
	if l.tail == nil {
		node.Next = node
		l.tail = node
	} else {
		node.Next = l.tail.Next
		l.tail.Next = node
		l.tail = node
	}
	l.size++
}

// Head returns the first node (tail.Next), or nil if the list is empty.
func (l *LinkedList[T]) Head() *Node[T] {
	if l.tail == nil {
		return nil
	}
	return l.tail.Next
}

// Size returns the number of elements in the list.
func (l *LinkedList[T]) Size() int {
	return l.size
}

// IsEmpty returns true if the list has no elements.
func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// ToSlice returns the list elements as a slice.
func (l *LinkedList[T]) ToSlice() []T {
	result := make([]T, 0, l.size)
	if l.tail == nil {
		return result
	}
	current := l.tail.Next
	for i := 0; i < l.size; i++ {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
