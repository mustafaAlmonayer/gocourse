package intermediate

import (
	"fmt"
)

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Stack[T any] struct {
	Elements []T
}

func (s *Stack[T]) push(element T) {
	s.Elements = append(s.Elements, element)
}

func (s *Stack[T]) pop() (T, bool) {
	var zero T
	if len(s.Elements) == 0 {
		return zero, false
	}
	v := s.Elements[len(s.Elements)-1]
	s.Elements[len(s.Elements)-1] = zero
	s.Elements = s.Elements[:len(s.Elements)-1]
	return v, true
}

func generics() {
	x, y := 1, 2
	x, y = swap(x, y)
	fmt.Println(x, y)

	s := Stack[int]{
		Elements: make([]int, 0),
	}

	s.push(10)
	s.push(20)
	s.push(30)
	s.pop()
	fmt.Println(s)
}
