package sets

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elems ...T) Set[T] {
	s := Set[T]{}
	s.Add(elems...)
	return s
}

func (s *Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		(*s)[elem] = struct{}{}
	}
}

func (s *Set[T]) Remove(elems ...T) {
	for _, elem := range elems {
		delete(*s, elem)
	}
}

func (s *Set[T]) Contains(elem T) bool {
	_, ok := (*s)[elem]
	return ok
}

func (s *Set[T]) Len() int {
	return len(*s)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set[T]) Print() {
	ret := make([]T, s.Len())
	i := 0
	for k := range *s {
		ret[i] = k
		i++
	}
	fmt.Println(ret)
}
