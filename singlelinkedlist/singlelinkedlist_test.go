package singlelinkedlist

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
