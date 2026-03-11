package doublelinkedlist

import (
	"reflect"
	"testing"
)

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

func TestDeleteMiddle(t *testing.T) {
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

func TestDeleteTail(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ll.Append(2)

	ll.Delete(2)
	expected := []int{1}
	got := ll.ToSlice()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	if ll.Tail().Value != 1 {
		t.Fatalf("expected tail value 1, got %d", ll.Tail().Value)
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

func TestHeadAndTail(t *testing.T) {
	ll := New[int]()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	if ll.Head().Value != 10 {
		t.Fatalf("expected head 10, got %d", ll.Head().Value)
	}
	if ll.Tail().Value != 30 {
		t.Fatalf("expected tail 30, got %d", ll.Tail().Value)
	}
}

func TestPrevPointers(t *testing.T) {
	ll := New[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	node := ll.Tail()
	var reversed []int
	for node != nil {
		reversed = append(reversed, node.Value)
		node = node.Prev
	}
	expected := []int{3, 2, 1}
	if !reflect.DeepEqual(reversed, expected) {
		t.Fatalf("expected %v traversing backwards, got %v", expected, reversed)
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
