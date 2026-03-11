package heap

import "cmp"

// Heap represents a min-heap data structure using generics.
type Heap[T cmp.Ordered] struct {
	items []T
}

// New creates and returns a new empty min-Heap.
func New[T cmp.Ordered]() *Heap[T] {
	return &Heap[T]{
		items: make([]T, 0),
	}
}

// Push adds an element to the heap.
func (h *Heap[T]) Push(item T) {
	h.items = append(h.items, item)
	h.siftUp(len(h.items) - 1)
}

// Pop removes and returns the minimum element. Returns false if empty.
func (h *Heap[T]) Pop() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	min := h.items[0]
	last := len(h.items) - 1
	h.items[0] = h.items[last]
	h.items = h.items[:last]
	if len(h.items) > 0 {
		h.siftDown(0)
	}
	return min, true
}

// Peek returns the minimum element without removing it. Returns false if empty.
func (h *Heap[T]) Peek() (T, bool) {
	if len(h.items) == 0 {
		var zero T
		return zero, false
	}
	return h.items[0], true
}

// Size returns the number of elements in the heap.
func (h *Heap[T]) Size() int {
	return len(h.items)
}

// IsEmpty returns true if the heap has no elements.
func (h *Heap[T]) IsEmpty() bool {
	return len(h.items) == 0
}

func (h *Heap[T]) siftUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.items[index] < h.items[parent] {
			h.items[index], h.items[parent] = h.items[parent], h.items[index]
			index = parent
		} else {
			break
		}
	}
}

func (h *Heap[T]) siftDown(index int) {
	size := len(h.items)
	for {
		smallest := index
		left := 2*index + 1
		right := 2*index + 2

		if left < size && h.items[left] < h.items[smallest] {
			smallest = left
		}
		if right < size && h.items[right] < h.items[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.items[index], h.items[smallest] = h.items[smallest], h.items[index]
		index = smallest
	}
}
