package deque

// Deque represents a double-ended queue using generics.
type Deque[T any] struct {
	items []T
}

// New creates and returns a new empty Deque.
func New[T any]() *Deque[T] {
	return &Deque[T]{
		items: make([]T, 0),
	}
}

// PushFront adds an element to the front of the deque.
func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

// PushBack adds an element to the back of the deque.
func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

// PopFront removes and returns the front element. Returns false if empty.
func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

// PopBack removes and returns the back element. Returns false if empty.
func (d *Deque[T]) PopBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item, true
}

// PeekFront returns the front element without removing it. Returns false if empty.
func (d *Deque[T]) PeekFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	return d.items[0], true
}

// PeekBack returns the back element without removing it. Returns false if empty.
func (d *Deque[T]) PeekBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	return d.items[len(d.items)-1], true
}

// Size returns the number of elements in the deque.
func (d *Deque[T]) Size() int {
	return len(d.items)
}

// IsEmpty returns true if the deque has no elements.
func (d *Deque[T]) IsEmpty() bool {
	return len(d.items) == 0
}
