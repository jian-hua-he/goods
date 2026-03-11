package tree

import "cmp"

// Node represents a single node in the binary search tree.
type Node[T cmp.Ordered] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

// BST represents a binary search tree.
type BST[T cmp.Ordered] struct {
	root *Node[T]
	size int
}

// New creates and returns a new empty BST.
func New[T cmp.Ordered]() *BST[T] {
	return &BST[T]{}
}

// Insert adds a value to the tree.
func (t *BST[T]) Insert(value T) {
	t.root = insertNode(t.root, value)
	t.size++
}

func insertNode[T cmp.Ordered](node *Node[T], value T) *Node[T] {
	if node == nil {
		return &Node[T]{Value: value}
	}
	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	}
	return node
}

// Search returns true if the value exists in the tree.
func (t *BST[T]) Search(value T) bool {
	return searchNode(t.root, value)
}

func searchNode[T cmp.Ordered](node *Node[T], value T) bool {
	if node == nil {
		return false
	}
	if value == node.Value {
		return true
	}
	if value < node.Value {
		return searchNode(node.Left, value)
	}
	return searchNode(node.Right, value)
}

// Delete removes a value from the tree. Returns false if not found.
func (t *BST[T]) Delete(value T) bool {
	var found bool
	t.root, found = deleteNode(t.root, value)
	if found {
		t.size--
	}
	return found
}

func deleteNode[T cmp.Ordered](node *Node[T], value T) (*Node[T], bool) {
	if node == nil {
		return nil, false
	}
	var found bool
	if value < node.Value {
		node.Left, found = deleteNode(node.Left, value)
	} else if value > node.Value {
		node.Right, found = deleteNode(node.Right, value)
	} else {
		// Found the node to delete
		if node.Left == nil {
			return node.Right, true
		}
		if node.Right == nil {
			return node.Left, true
		}
		// Node has two children: find in-order successor (min of right subtree)
		successor := findMin(node.Right)
		node.Value = successor.Value
		node.Right, _ = deleteNode(node.Right, successor.Value)
		return node, true
	}
	return node, found
}

func findMin[T cmp.Ordered](node *Node[T]) *Node[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Min returns the minimum value in the tree. Returns false if empty.
func (t *BST[T]) Min() (T, bool) {
	if t.root == nil {
		var zero T
		return zero, false
	}
	return findMin(t.root).Value, true
}

// Max returns the maximum value in the tree. Returns false if empty.
func (t *BST[T]) Max() (T, bool) {
	if t.root == nil {
		var zero T
		return zero, false
	}
	current := t.root
	for current.Right != nil {
		current = current.Right
	}
	return current.Value, true
}

// InOrder returns values in sorted order (in-order traversal).
func (t *BST[T]) InOrder() []T {
	result := make([]T, 0, t.size)
	inOrderTraversal(t.root, &result)
	return result
}

func inOrderTraversal[T cmp.Ordered](node *Node[T], result *[]T) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left, result)
	*result = append(*result, node.Value)
	inOrderTraversal(node.Right, result)
}

// Size returns the number of elements in the tree.
func (t *BST[T]) Size() int {
	return t.size
}

// IsEmpty returns true if the tree has no elements.
func (t *BST[T]) IsEmpty() bool {
	return t.size == 0
}
