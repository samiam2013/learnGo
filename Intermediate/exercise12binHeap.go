package main

import "fmt"

// Implement a binary heap by implementing a pointer-linked binary tree.
// Use it for implementing heap-sort.

type binHeap struct {
	root *binNode
}

type binNode struct {
	value       int
	parent      *binNode
	left, right *binNode
}

//log.Printf("var: %#+v\n", var)
// it's not specified so I'll implement a max-heap of numbers
func (b *binHeap) add(newVal int) {
	// heapify
	// put the value in the left-most bottom growth
	if b.root.value == 0 && (b.root.left == nil && b.root.right == nil) {
		b.root.value = newVal
		return
	}
	toVisit := []*binNode{b.root}
	current := toVisit[0]
	lastParent := current
	for i := 1; current != nil; i++ {
		toVisit = append(toVisit, current.left, current.right)
		lastParent = current
		current = toVisit[i]
	}
	current = &binNode{newVal, lastParent, nil, nil}
	// Compare the added element with its parent; if they are in the correct order, stop.
	//  if not, swap the element with its parent and return to the previous step.
}

func exercise12() {
	fmt.Println("exercise 12...")
	heap := binHeap{&binNode{0, nil, nil, nil}}
	for i := 0; i < 1000; i++ {
		heap.add(i)
	}
	fmt.Println(heap.root.value)
}
