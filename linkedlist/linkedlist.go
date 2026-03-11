package linkedlist

// Node represents a single node in the linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a singly linked list.
type LinkedList[T comparable] struct {
	head *Node[T]
	size int
}

// New creates and returns a new empty LinkedList.
func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Prepend adds an element to the front of the list.
func (l *LinkedList[T]) Prepend(value T) {
	node := &Node[T]{Value: value, Next: l.head}
	l.head = node
	l.size++
}

// Append adds an element to the end of the list.
func (l *LinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value}
	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
	l.size++
}

// Delete removes the first occurrence of the value. Returns false if not found.
func (l *LinkedList[T]) Delete(value T) bool {
	if l.head == nil {
		return false
	}
	if l.head.Value == value {
		l.head = l.head.Next
		l.size--
		return true
	}
	current := l.head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			l.size--
			return true
		}
		current = current.Next
	}
	return false
}

// Head returns the first node, or nil if the list is empty.
func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
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
