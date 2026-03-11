package queue

import "testing"

func TestNew(t *testing.T) {
	q := New[int]()
	if q == nil {
		t.Fatal("expected non-nil queue")
	}
}

func TestEnqueueAndDequeue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Fatalf("expected 1, got %d (ok=%v)", val, ok)
	}
	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Fatalf("expected 2, got %d (ok=%v)", val, ok)
	}
	val, ok = q.Dequeue()
	if !ok || val != 3 {
		t.Fatalf("expected 3, got %d (ok=%v)", val, ok)
	}
}

func TestDequeueEmpty(t *testing.T) {
	q := New[string]()
	_, ok := q.Dequeue()
	if ok {
		t.Fatal("expected ok=false on empty queue")
	}
}

func TestPeek(t *testing.T) {
	q := New[int]()
	q.Enqueue(42)

	val, ok := q.Peek()
	if !ok || val != 42 {
		t.Fatalf("expected 42, got %d (ok=%v)", val, ok)
	}
	// Peek should not remove the element
	if q.Size() != 1 {
		t.Fatalf("expected size 1 after peek, got %d", q.Size())
	}
}

func TestPeekEmpty(t *testing.T) {
	q := New[int]()
	_, ok := q.Peek()
	if ok {
		t.Fatal("expected ok=false on empty queue")
	}
}

func TestSize(t *testing.T) {
	q := New[int]()
	if q.Size() != 0 {
		t.Fatalf("expected size 0, got %d", q.Size())
	}
	q.Enqueue(1)
	q.Enqueue(2)
	if q.Size() != 2 {
		t.Fatalf("expected size 2, got %d", q.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	q := New[int]()
	if !q.IsEmpty() {
		t.Fatal("expected empty queue")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Fatal("expected non-empty queue")
	}
}

func TestStringQueue(t *testing.T) {
	q := New[string]()
	q.Enqueue("hello")
	q.Enqueue("world")

	val, ok := q.Dequeue()
	if !ok || val != "hello" {
		t.Fatalf("expected hello, got %s", val)
	}
}
