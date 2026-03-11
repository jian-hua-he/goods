package heap

import "testing"

func TestNew(t *testing.T) {
	h := New[int]()
	if h == nil {
		t.Fatal("expected non-nil heap")
	}
}

func TestPushAndPeek(t *testing.T) {
	h := New[int]()
	h.Push(5)
	h.Push(3)
	h.Push(7)

	val, ok := h.Peek()
	if !ok || val != 3 {
		t.Fatalf("expected min=3, got %d (ok=%v)", val, ok)
	}
}

func TestPop(t *testing.T) {
	h := New[int]()
	h.Push(5)
	h.Push(3)
	h.Push(7)
	h.Push(1)

	val, ok := h.Pop()
	if !ok || val != 1 {
		t.Fatalf("expected 1, got %d (ok=%v)", val, ok)
	}
	val, ok = h.Pop()
	if !ok || val != 3 {
		t.Fatalf("expected 3, got %d (ok=%v)", val, ok)
	}
	val, ok = h.Pop()
	if !ok || val != 5 {
		t.Fatalf("expected 5, got %d (ok=%v)", val, ok)
	}
	val, ok = h.Pop()
	if !ok || val != 7 {
		t.Fatalf("expected 7, got %d (ok=%v)", val, ok)
	}
}

func TestPopEmpty(t *testing.T) {
	h := New[int]()
	_, ok := h.Pop()
	if ok {
		t.Fatal("expected ok=false on empty heap")
	}
}

func TestPeekEmpty(t *testing.T) {
	h := New[int]()
	_, ok := h.Peek()
	if ok {
		t.Fatal("expected ok=false on empty heap")
	}
}

func TestSize(t *testing.T) {
	h := New[int]()
	if h.Size() != 0 {
		t.Fatalf("expected size 0, got %d", h.Size())
	}
	h.Push(1)
	h.Push(2)
	if h.Size() != 2 {
		t.Fatalf("expected size 2, got %d", h.Size())
	}
	h.Pop()
	if h.Size() != 1 {
		t.Fatalf("expected size 1, got %d", h.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	h := New[int]()
	if !h.IsEmpty() {
		t.Fatal("expected empty heap")
	}
	h.Push(1)
	if h.IsEmpty() {
		t.Fatal("expected non-empty heap")
	}
}

func TestHeapProperty(t *testing.T) {
	h := New[int]()
	values := []int{10, 4, 15, 1, 7, 20, 3}
	for _, v := range values {
		h.Push(v)
	}

	prev := -1
	for !h.IsEmpty() {
		val, _ := h.Pop()
		if val < prev {
			t.Fatalf("heap property violated: got %d after %d", val, prev)
		}
		prev = val
	}
}

func TestStringHeap(t *testing.T) {
	h := New[string]()
	h.Push("banana")
	h.Push("apple")
	h.Push("cherry")

	val, ok := h.Pop()
	if !ok || val != "apple" {
		t.Fatalf("expected apple, got %s", val)
	}
}
