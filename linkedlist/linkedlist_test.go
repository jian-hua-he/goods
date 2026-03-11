package linkedlist

import (
	"reflect"
	"testing"
)

func TestPrepend(t *testing.T) {
	ll := New[int]()
	ll.Prepend(3)
	ll.Prepend(2)
	ll.Prepend(1)

	expected := []int{1, 2, 3}
	got := ll.ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestAppend(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	expected := []int{1, 2, 3}
	got := ll.ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestDelete(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	ok := ll.Delete(2)
	if !ok {
		t.Fatal("expected delete to return true")
	}
	expected := []int{1, 3}
	got := ll.ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestDeleteHead(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ll.Append(2)

	ll.Delete(1)
	expected := []int{2}
	got := ll.ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestDeleteNotFound(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ok := ll.Delete(99)
	if ok {
		t.Fatal("expected delete to return false for missing value")
	}
}

func TestSearch(t *testing.T) {
	ll := New[string]()
	ll.Append("hello")
	ll.Append("world")

	if !ll.Search("world") {
		t.Fatal("expected to find 'world'")
	}
	if ll.Search("missing") {
		t.Fatal("expected not to find 'missing'")
	}
}

func TestGet(t *testing.T) {
	ll := New[int]()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	val, ok := ll.Get(1)
	if !ok || val != 20 {
		t.Fatalf("expected 20, got %d (ok=%v)", val, ok)
	}

	_, ok = ll.Get(5)
	if ok {
		t.Fatal("expected ok=false for out of bounds")
	}
}

func TestSize(t *testing.T) {
	ll := New[int]()
	if ll.Size() != 0 {
		t.Fatalf("expected size 0, got %d", ll.Size())
	}
	ll.Append(1)
	ll.Append(2)
	if ll.Size() != 2 {
		t.Fatalf("expected size 2, got %d", ll.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	ll := New[int]()
	if !ll.IsEmpty() {
		t.Fatal("expected empty list")
	}
	ll.Append(1)
	if ll.IsEmpty() {
		t.Fatal("expected non-empty list")
	}
}

func TestToSliceEmpty(t *testing.T) {
	ll := New[int]()
	got := ll.ToSlice()
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}
