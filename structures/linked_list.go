package structures

import (
	"fmt"
	"strings"
)

/* A node within the linked list containing the data and pointer to next node */
type LinkedListNode[T any] struct {
	Data T
	Next *LinkedListNode[T]
}

/* The main linked list structure with all associated methods */
type LinkedList[T any] struct {
	length int
	Head   *LinkedListNode[T]
	Tail   *LinkedListNode[T]
}

/* Returns an empty LinkedList */
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{length: 0}
}

/* Checks the LinkedList index boundaries for legal operations. */
/* Returns the index on success, or an error message otherwise. */
/* Is capable of handling a negative index up to the negative length of the LinkedList. */
func (ll *LinkedList[T]) BoundaryCheck(index int) (int, error) {
	newIndex := index
	if index < 0 {
		newIndex = ll.length + index
	}

	if ll.length < 1 || newIndex < 0 || newIndex > ll.length-1 {
		return -1, fmt.Errorf("index %d out of bounds for LinkedList of length %d", index, ll.length)
	}

	return newIndex, nil
}

/* Retrives the value at a given index on success, returns an error otherwise */
func (ll *LinkedList[T]) GetNode(index int) (*LinkedListNode[T], error) {
	index, err := ll.BoundaryCheck(index)
	if err != nil {
		return nil, err
	}

	node := ll.Head
	for i := 0; i <= index; i++ {
		if i == index {
			break
		}
		node = node.Next
	}

	return node, nil
}

/* Appends data to the end of the list */
func (ll *LinkedList[T]) Append(data T) {
	node := &LinkedListNode[T]{Data: data, Next: nil}

	if ll.length < 1 {
		ll.Head = node
	} else {
		ll.Tail.Next = node
	}

	ll.Tail = node
	ll.length++
}

/* Inserts data at the beginning of the list */
func (ll *LinkedList[T]) Insert(data T) {
	next := ll.Head
	ll.Head = &LinkedListNode[T]{Data: data, Next: next}
	ll.length++
}

/* Inserts data at a given index on success, otherwise returns an error */
func (ll *LinkedList[T]) InsertAt(index int, data T) error {
	previousNode, err := ll.GetNode(index - 1)
	if err != nil {
		return err
	}

	node := &LinkedListNode[T]{Data: data, Next: previousNode.Next}
	previousNode.Next = node
	ll.length++
	return nil
}

/* Removes and returns the value of the first element in the list, returns an error otherwise */
func (ll *LinkedList[T]) Shift() (*T, error) {
	if ll.length < 1 {
		return nil, fmt.Errorf("cannot shift an empty list")
	}

	data := &ll.Head.Data

	if ll.length == 1 {
		ll.Head = nil
		ll.Tail = nil
		return data, nil
	}

	ll.Head = ll.Head.Next
	ll.length--

	return data, nil
}

/* Removes and returns the value of the last element in the list, returns an error otherwise */
func (ll *LinkedList[T]) Pop() (*T, error) {
	if ll.length < 1 {
		return nil, fmt.Errorf("cannot pop an empty list")
	}

	var data *T
	if ll.length == 1 {
		data = &ll.Head.Data
		ll.Head = nil
		ll.Tail = nil
	} else {
		node := ll.Head
		for node.Next != ll.Tail {
			node = node.Next
		}

		data = &ll.Tail.Data
		node.Next = nil
		ll.Tail = node
	}

	ll.length--
	return data, nil
}

/* Removes and returns the value of the element at the given index, returns an error otherwise */
func (ll *LinkedList[T]) PopAt(index int) (*T, error) {
	node, err := ll.GetNode(index - 1)
	if err != nil {
		return nil, err
	}

	data := node.Next.Data
	node.Next = node.Next.Next
	ll.length--

	return &data, nil
}

/* Prints the contents of the entire list */
func (ll *LinkedList[T]) Print() {
	var buf strings.Builder

	if ll.length < 1 {
		fmt.Println("{ }")
		return
	}

	next := ll.Head
	for next != nil {
		fmt.Fprintf(&buf, "%v ", next.Data)
		next = next.Next
	}

	fmt.Printf("{ %s}\n", buf.String())
}

/* Prints the contents at the given index, prints an error otherwise */
func (ll *LinkedList[T]) PrintAt(index int) {
	if node, err := ll.GetNode(index); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(node.Data)
	}
}

/* Returns the length of the list */
func (ll *LinkedList[T]) Length() int {
	return ll.length
}
