package stack

import "testing"

func TestNew(t *testing.T) {
	s := New[int]()
	if s == nil {
		t.Fatal("expected non-nil stack")
	}
}

func TestPushAndPop(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	val, ok := s.Pop()
	if !ok || val != 3 {
		t.Fatalf("expected 3, got %d (ok=%v)", val, ok)
	}
	val, ok = s.Pop()
	if !ok || val != 2 {
		t.Fatalf("expected 2, got %d (ok=%v)", val, ok)
	}
	val, ok = s.Pop()
	if !ok || val != 1 {
		t.Fatalf("expected 1, got %d (ok=%v)", val, ok)
	}
}

func TestPopEmpty(t *testing.T) {
	s := New[string]()
	_, ok := s.Pop()
	if ok {
		t.Fatal("expected ok=false on empty stack")
	}
}

func TestPeek(t *testing.T) {
	s := New[int]()
	s.Push(42)

	val, ok := s.Peek()
	if !ok || val != 42 {
		t.Fatalf("expected 42, got %d (ok=%v)", val, ok)
	}
	if s.Size() != 1 {
		t.Fatalf("expected size 1 after peek, got %d", s.Size())
	}
}

func TestPeekEmpty(t *testing.T) {
	s := New[int]()
	_, ok := s.Peek()
	if ok {
		t.Fatal("expected ok=false on empty stack")
	}
}

func TestSize(t *testing.T) {
	s := New[int]()
	if s.Size() != 0 {
		t.Fatalf("expected size 0, got %d", s.Size())
	}
	s.Push(1)
	s.Push(2)
	if s.Size() != 2 {
		t.Fatalf("expected size 2, got %d", s.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	s := New[int]()
	if !s.IsEmpty() {
		t.Fatal("expected empty stack")
	}
	s.Push(1)
	if s.IsEmpty() {
		t.Fatal("expected non-empty stack")
	}
}

func TestLIFOOrder(t *testing.T) {
	s := New[string]()
	s.Push("a")
	s.Push("b")
	s.Push("c")

	val, _ := s.Pop()
	if val != "c" {
		t.Fatalf("expected c, got %s", val)
	}
	val, _ = s.Pop()
	if val != "b" {
		t.Fatalf("expected b, got %s", val)
	}
}
