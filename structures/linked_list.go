package structures

import (
	"fmt"
	"strings"
)

type linkedListNode[T any] struct {
	Value T
	Next  *linkedListNode[T]
	Prev  *linkedListNode[T]
}

type linkedList[T any] struct {
	length int
	head   *linkedListNode[T]
	tail   *linkedListNode[T]
}

func NewLinkedList[T any]() *linkedList[T] {
	return &linkedList[T]{}
}

// checks the index is within bounds, accounts for negative indexing
// returns positive index on success or -1 on error
func (ll *linkedList[T]) boundaryCheck(index int) int {
	newIndex := index
	if index < 0 {
		newIndex = ll.length + index
	}

	if ll.length > 0 && newIndex >= 0 && newIndex < ll.length {
		return newIndex
	}

	panic(fmt.Sprintf("index %d out of bounds for list of length %d\n", index, ll.length))
}

// returns the node at the given index or panics on error
func (ll *linkedList[T]) getNode(index int) *linkedListNode[T] {
	index = ll.boundaryCheck(index)

	if index == 0 {
		return ll.head
	} else if index == ll.length-1 {
		return ll.tail
	}

	var node *linkedListNode[T]
	if min(index, ll.length-index) != index {
		node = ll.tail
		for i := ll.length - 1; i >= index; i-- {
			if i == index {
				break
			}
			node = node.Prev
		}
	} else {
		node = ll.head
		for i := 0; i <= index; i++ {
			if i == index {
				break
			}
			node = node.Next
		}
	}

	return node
}

func (ll *linkedList[T]) Length() int {
	return ll.length
}

// returns the value at the given index
func (ll *linkedList[T]) ValueAt(index int) T {
	return ll.getNode(index).Value
}

// appends data to the end of the list
func (ll *linkedList[T]) Append(data T) {
	node := &linkedListNode[T]{Value: data}

	if ll.length == 0 {
		ll.head = node
		ll.tail = ll.head
	} else {
		node.Prev = ll.tail
		ll.tail.Next = node
		ll.tail = node
	}

	ll.length++
}

// inserts data at the specified index or panics on error
func (ll *linkedList[T]) Insert(data T, index int) {
	if index == 0 && ll.length == 0 {
		ll.Append(data)
		return
	}

	node := &linkedListNode[T]{Value: data}
	if index == 0 || index == -ll.length {
		node.Next = ll.head
		ll.head.Prev = node
		ll.head = node
	} else {
		current := ll.getNode(index)
		node.Next = current
		node.Prev = current.Prev
		current.Prev.Next = node
		current.Prev = node
	}

	ll.length++
}

// removes and returns the value from the linked list at the given index or panics on error
func (ll *linkedList[T]) Pop(index int) T {
	var value T

	if ll.length == 1 {
		value = ll.head.Value
		ll.head = nil
		ll.tail = ll.head
	} else if index == 0 || index == -ll.length {
		value = ll.head.Value
		ll.head.Next.Prev = nil
		ll.head = ll.head.Next
	} else if index == ll.length-1 || index == -1 {
		value = ll.tail.Value
		ll.tail.Prev.Next = nil
		ll.tail = ll.tail.Prev
	} else {
		node := ll.getNode(index)
		value = node.Value
		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
	}

	ll.length--
	return value
}

// clears all items out of the linked list
func (ll *linkedList[T]) Clear() {
	ll.head = nil
	ll.tail = nil
}

// Prints the contents of the entire list
func (ll *linkedList[T]) Print() {
	var buf strings.Builder

	if ll.length < 1 {
		fmt.Println("{ }")
		return
	}

	next := ll.head
	for next != nil {
		fmt.Fprintf(&buf, "%v ", next.Value)
		next = next.Next
	}

	fmt.Printf("{ %s}\n", buf.String())
}
