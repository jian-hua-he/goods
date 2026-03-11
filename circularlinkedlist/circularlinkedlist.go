package circularlinkedlist

// Node represents a single node in a circular linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// LinkedList represents a circular singly linked list.
type LinkedList[T comparable] struct {
	tail *Node[T]
	size int
}

// New creates and returns a new empty LinkedList.
func New[T comparable]() *LinkedList[T] {
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

// Delete removes the first occurrence of the value. Returns false if not found.
func (l *LinkedList[T]) Delete(value T) bool {
	if l.tail == nil {
		return false
	}
	// Single element
	if l.size == 1 {
		if l.tail.Value == value {
			l.tail = nil
			l.size--
			return true
		}
		return false
	}
	// Check if head (tail.Next) matches
	head := l.tail.Next
	if head.Value == value {
		l.tail.Next = head.Next
		l.size--
		return true
	}
	// Traverse the rest
	current := head
	for current.Next != l.tail.Next {
		if current.Next.Value == value {
			if current.Next == l.tail {
				l.tail = current
			}
			current.Next = current.Next.Next
			l.size--
			return true
		}
		current = current.Next
	}
	return false
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
