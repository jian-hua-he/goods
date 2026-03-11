# goods

A Go library providing generic implementations of common data structures.

## Data Structures

| Package | Structure | Description |
|---------|-----------|-------------|
| `queue` | Queue | FIFO queue with `Enqueue`, `Dequeue`, `Peek`, `Size`, `IsEmpty` |
| `stack` | Stack | LIFO stack with `Push`, `Pop`, `Peek`, `Size`, `IsEmpty` |
| `linkedlist` | Linked List | Singly linked list with `Prepend`, `Append`, `Delete`, `Search`, `Get`, `Size`, `IsEmpty`, `ToSlice` |
| `tree` | Binary Search Tree | BST with `Insert`, `Search`, `Delete`, `Min`, `Max`, `InOrder`, `Size`, `IsEmpty` |
| `heap` | Min-Heap | Min-heap with `Push`, `Pop`, `Peek`, `Size`, `IsEmpty` |

## Requirements

- Go 1.22+

## Installation

```sh
go get github.com/jian-hua-he/goods
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/jian-hua-he/goods/queue"
	"github.com/jian-hua-he/goods/stack"
	"github.com/jian-hua-he/goods/linkedlist"
	"github.com/jian-hua-he/goods/tree"
	"github.com/jian-hua-he/goods/heap"
)

func main() {
	// Queue
	q := queue.New[string]()
	q.Enqueue("first")
	q.Enqueue("second")
	val, _ := q.Dequeue()
	fmt.Println(val) // "first"

	// Stack
	s := stack.New[int]()
	s.Push(1)
	s.Push(2)
	val2, _ := s.Pop()
	fmt.Println(val2) // 2

	// Linked List
	ll := linkedlist.New[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Prepend(0)
	fmt.Println(ll.ToSlice()) // [0 1 2]

	// Binary Search Tree
	bst := tree.New[int]()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(7)
	fmt.Println(bst.InOrder()) // [3 5 7]

	// Min-Heap
	h := heap.New[int]()
	h.Push(5)
	h.Push(1)
	h.Push(3)
	val3, _ := h.Pop()
	fmt.Println(val3) // 1
}
```

## Testing

```sh
go test ./...
```
