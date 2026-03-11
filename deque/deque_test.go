package deque

import "testing"

func TestPushFrontPopFront(t *testing.T) {
	d := New[int]()
	d.PushFront(3)
	d.PushFront(2)
	d.PushFront(1)

	for i := 1; i <= 3; i++ {
		val, ok := d.PopFront()
		if !ok || val != i {
			t.Fatalf("expected %d, got %d (ok=%v)", i, val, ok)
		}
	}
}

func TestPushBackPopBack(t *testing.T) {
	d := New[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	for i := 3; i >= 1; i-- {
		val, ok := d.PopBack()
		if !ok || val != i {
			t.Fatalf("expected %d, got %d (ok=%v)", i, val, ok)
		}
	}
}

func TestPushFrontPopBack(t *testing.T) {
	d := New[string]()
	d.PushFront("a")
	d.PushFront("b")

	val, ok := d.PopBack()
	if !ok || val != "a" {
		t.Fatalf("expected 'a', got '%s'", val)
	}
	val, ok = d.PopBack()
	if !ok || val != "b" {
		t.Fatalf("expected 'b', got '%s'", val)
	}
}

func TestPushBackPopFront(t *testing.T) {
	d := New[int]()
	d.PushBack(10)
	d.PushBack(20)

	val, ok := d.PopFront()
	if !ok || val != 10 {
		t.Fatalf("expected 10, got %d", val)
	}
	val, ok = d.PopFront()
	if !ok || val != 20 {
		t.Fatalf("expected 20, got %d", val)
	}
}

func TestPopEmptyDeque(t *testing.T) {
	d := New[int]()

	_, ok := d.PopFront()
	if ok {
		t.Fatal("expected ok=false for PopFront on empty deque")
	}
	_, ok = d.PopBack()
	if ok {
		t.Fatal("expected ok=false for PopBack on empty deque")
	}
}

func TestPeekFront(t *testing.T) {
	d := New[int]()
	d.PushBack(1)
	d.PushBack(2)

	val, ok := d.PeekFront()
	if !ok || val != 1 {
		t.Fatalf("expected 1, got %d", val)
	}
	if d.Size() != 2 {
		t.Fatal("PeekFront should not remove elements")
	}
}

func TestPeekBack(t *testing.T) {
	d := New[int]()
	d.PushBack(1)
	d.PushBack(2)

	val, ok := d.PeekBack()
	if !ok || val != 2 {
		t.Fatalf("expected 2, got %d", val)
	}
	if d.Size() != 2 {
		t.Fatal("PeekBack should not remove elements")
	}
}

func TestPeekEmpty(t *testing.T) {
	d := New[int]()

	_, ok := d.PeekFront()
	if ok {
		t.Fatal("expected ok=false for PeekFront on empty deque")
	}
	_, ok = d.PeekBack()
	if ok {
		t.Fatal("expected ok=false for PeekBack on empty deque")
	}
}

func TestSize(t *testing.T) {
	d := New[int]()
	if d.Size() != 0 {
		t.Fatalf("expected size 0, got %d", d.Size())
	}
	d.PushBack(1)
	d.PushFront(2)
	if d.Size() != 2 {
		t.Fatalf("expected size 2, got %d", d.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	d := New[int]()
	if !d.IsEmpty() {
		t.Fatal("expected empty deque")
	}
	d.PushBack(1)
	if d.IsEmpty() {
		t.Fatal("expected non-empty deque")
	}
}
