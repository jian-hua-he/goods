package doublelinkedlist

// Node represents a single node in a doubly linked list.
type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// LinkedList represents a doubly linked list.
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// New creates and returns a new empty LinkedList.
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append adds an element to the end of the list.
func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value, Prev: l.tail}
	if l.tail != nil {
		l.tail.Next = node
	} else {
		l.head = node
	}
	l.tail = node
	l.size++
}

// Prepend adds an element to the front of the list.
func (l *LinkedList[T]) Prepend(value T) {
	node := &Node[T]{Value: value, Next: l.head}
	if l.head != nil {
		l.head.Prev = node
	} else {
		l.tail = node
	}
	l.head = node
	l.size++
}

// Head returns the first node, or nil if the list is empty.
func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

// Tail returns the last node, or nil if the list is empty.
func (l *LinkedList[T]) Tail() *Node[T] {
	return l.tail
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
	current := l.head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
