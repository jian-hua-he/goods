package queue

// Queue represents a FIFO data structure using generics.
type Queue[T any] struct {
	items []T
}

// New creates and returns a new empty Queue.
func New[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the front element. Returns false if empty.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// Peek returns the front element without removing it. Returns false if empty.
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	return q.items[0], true
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// IsEmpty returns true if the queue has no elements.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
