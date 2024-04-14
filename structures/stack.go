package structures

import (
	"fmt"
	"strings"
)

type stack[T any] struct {
	static bool
	size   int
	top    int
	data   []*T
}

// initialises and returns a dynamically sized stack
func DynamicStack[T any]() *stack[T] {
	return &stack[T]{
		static: false,
		size:   0,
		top:    -1,
		data:   make([]*T, 0),
	}
}

// initialises and returns a fixed sized stack
func StaticStack[T any](size int) *stack[T] {
	return &stack[T]{
		static: true,
		size:   size,
		top:    -1,
		data:   make([]*T, size),
	}
}

// adds data to the top of the stack
func (s *stack[T]) Push(data T) error {
	switch s.static {
	case true:
		if s.top == s.size-1 {
			return fmt.Errorf("stack overflow")
		} else {
			s.data[s.top+1] = &data
		}

	default:
		s.data = append(s.data, &data)
		s.size = len(s.data)
	}

	s.top++
	return nil
}

// removes data from the top of the stack
func (s *stack[T]) Pop() (T, error) {
	var value T
	if s.top < 0 {
		return value, fmt.Errorf("empty stack")
	}

	value = *s.data[s.top]
	if s.static {
		s.data[s.top] = nil
	} else {
		s.data = append([]*T(nil), s.data[:s.top]...)
		s.size = len(s.data)
	}

	s.top--
	return value, nil
}

// Prints the contents of the entire stack
func (s *stack[T]) Print() {
	var buf strings.Builder

	if s.top < 0 {
		fmt.Println("{ }")
		return
	}

	for i := 0; i <= s.top; i++ {
		fmt.Fprintf(&buf, "%+v ", *s.data[i])
	}

	fmt.Printf("{ %s}\n", buf.String())
}
