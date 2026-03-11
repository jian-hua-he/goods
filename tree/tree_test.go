package tree

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	bst := New[int]()
	if bst == nil {
		t.Fatal("expected non-nil BST")
	}
}

func TestInsertAndSearch(t *testing.T) {
	bst := New[int]()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)

	if !bst.Search(5) {
		t.Fatal("expected to find 5")
	}
	if !bst.Search(3) {
		t.Fatal("expected to find 3")
	}
	if !bst.Search(7) {
		t.Fatal("expected to find 7")
	}
	if bst.Search(99) {
		t.Fatal("expected not to find 99")
	}
}

func TestInOrder(t *testing.T) {
	bst := New[int]()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(4)

	expected := []int{1, 3, 4, 5, 7}
	got := bst.InOrder()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestDelete(t *testing.T) {
	bst := New[int]()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)
	bst.Insert(4)

	// Delete leaf node
	ok := bst.Delete(1)
	if !ok {
		t.Fatal("expected delete to return true")
	}
	if bst.Search(1) {
		t.Fatal("expected 1 to be deleted")
	}

	// Delete node with one child
	bst.Delete(3)
	expected := []int{4, 5, 7}
	got := bst.InOrder()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}

	// Delete node with two children
	bst.Insert(6)
	bst.Insert(8)
	bst.Delete(7)
	expected = []int{4, 5, 6, 8}
	got = bst.InOrder()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestDeleteNotFound(t *testing.T) {
	bst := New[int]()
	bst.Insert(5)
	ok := bst.Delete(99)
	if ok {
		t.Fatal("expected delete to return false for missing value")
	}
}

func TestDeleteRoot(t *testing.T) {
	bst := New[int]()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)

	bst.Delete(5)
	if bst.Search(5) {
		t.Fatal("expected root to be deleted")
	}
	if bst.Size() != 2 {
		t.Fatalf("expected size 2, got %d", bst.Size())
	}
}

func TestMin(t *testing.T) {
	bst := New[int]()
	_, ok := bst.Min()
	if ok {
		t.Fatal("expected ok=false on empty tree")
	}

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(1)

	val, ok := bst.Min()
	if !ok || val != 1 {
		t.Fatalf("expected 1, got %d (ok=%v)", val, ok)
	}
}

func TestMax(t *testing.T) {
	bst := New[int]()
	_, ok := bst.Max()
	if ok {
		t.Fatal("expected ok=false on empty tree")
	}

	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	bst.Insert(9)

	val, ok := bst.Max()
	if !ok || val != 9 {
		t.Fatalf("expected 9, got %d (ok=%v)", val, ok)
	}
}

func TestSize(t *testing.T) {
	bst := New[int]()
	if bst.Size() != 0 {
		t.Fatalf("expected size 0, got %d", bst.Size())
	}
	bst.Insert(5)
	bst.Insert(3)
	if bst.Size() != 2 {
		t.Fatalf("expected size 2, got %d", bst.Size())
	}
}

func TestIsEmpty(t *testing.T) {
	bst := New[int]()
	if !bst.IsEmpty() {
		t.Fatal("expected empty tree")
	}
	bst.Insert(1)
	if bst.IsEmpty() {
		t.Fatal("expected non-empty tree")
	}
}

func TestStringBST(t *testing.T) {
	bst := New[string]()
	bst.Insert("banana")
	bst.Insert("apple")
	bst.Insert("cherry")

	expected := []string{"apple", "banana", "cherry"}
	got := bst.InOrder()
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
